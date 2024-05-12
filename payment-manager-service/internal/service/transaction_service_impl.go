package service

import (
	"context"
	"fmt"
	"net/http"
	"payment-manager-service/internal/delivery/http/exception"
	"payment-manager-service/internal/entity"
	"payment-manager-service/internal/model"
	"payment-manager-service/internal/repository"
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	DB                    *gorm.DB
	Validate              *validator.Validate
	Log                   *logrus.Logger
	UserRepository        *repository.UserRepository
	TransactionRepository *repository.TransactionRepository
	AccountRepository     *repository.AccountRepository
	CurrencyRepository    *repository.CurrencyRepository
}

func NewTransactionService(db *gorm.DB, validate *validator.Validate, log *logrus.Logger,
	userRepository *repository.UserRepository, transactionRepository *repository.TransactionRepository,
	accountRepository *repository.AccountRepository, currencyRepository *repository.CurrencyRepository) TransactionService {
	return &TransactionServiceImpl{
		DB:                    db,
		Validate:              validate,
		Log:                   log,
		UserRepository:        userRepository,
		AccountRepository:     accountRepository,
		TransactionRepository: transactionRepository,
	}
}

func (s *TransactionServiceImpl) Transfer(ctx context.Context, request *model.TransferRequest) (*model.TransferResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("[%d] invalid request body : %+v", http.StatusBadRequest, err)
		return nil, err
	}

	tx := s.DB.WithContext(ctx).Begin()

	senderAccount := new(entity.Account)
	senderAccount.ID = request.SenderAccountID
	err := s.AccountRepository.FindByIDWithAccountType(tx, senderAccount)
	if err != nil {
		s.Log.Warnf("[%d] sender account not found : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	senderUser := new(entity.User)
	if err := s.UserRepository.Repository.FindById(tx, senderUser, request.UserID); err != nil {
		s.Log.Warnf("[%d] invalid sender account id : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	if request.UserID != senderAccount.UserID {
		s.Log.Warnf("[%d] invalid sender account id : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	recipientAccount := new(entity.Account)
	recipientAccount.Number = request.RecipientAccountNumber
	err = s.AccountRepository.FindByAccountNumber(tx, recipientAccount)
	if err != nil {
		s.Log.Warnf("[%d] recipient account not found : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find recipient account")
	}

	recipientUser := new(entity.User)
	if err := s.UserRepository.Repository.FindById(tx, recipientUser, recipientAccount.UserID); err != nil {
		s.Log.Warnf("[%d] invalid sender account id : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	currency := new(entity.Currency)
	currency.Code = request.CurrencyCode
	err = s.CurrencyRepository.FindCurrentCurrencyByCode(tx, currency)
	if err != nil {
		s.Log.Warnf("[%d] currency not found : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find currency")
	}

	transferAmountInIDR := request.Amount / currency.CurrentPer1IDR
	if senderAccount.BalanceIDR < (transferAmountInIDR + senderAccount.AccountType.LimitIDR) {
		s.Log.Warnf("[%d] sender insufficient balance : %+v", http.StatusBadRequest, err)
		return nil, exception.NewHttpError(http.StatusBadRequest, "insufficient balance")
	}

	transaction := &entity.Transaction{
		ID:                 uuid.NewString(),
		SenderAccountID:    senderAccount.ID,
		RecipientAccountID: recipientAccount.ID,
		CurrencyID:         currency.ID,
		OriginalAmount:     request.Amount,
		IDRAmount:          transferAmountInIDR,
		Status:             1, // Set as pending
		CreatedAt:          time.Now(),
	}
	if err := s.TransactionRepository.Create(tx, transaction); err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			s.Log.Warnf("[%d] %+v", http.StatusInternalServerError, errRollback)
			return nil, exception.NewHttpError(http.StatusInternalServerError, "something wrong")
		}

		s.Log.Warnf("[%d] %+v", http.StatusInternalServerError, err)
		return nil, exception.NewHttpError(http.StatusInternalServerError, "something wrong")
	}
	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("[%d] %+v", http.StatusInternalServerError, err)
		return nil, exception.NewHttpError(http.StatusInternalServerError, "something wrong")
	}

	// Set automatically succeed, if transaction hasn't been cancelled
	time.AfterFunc(30*time.Second, func() {
		tx := s.DB.Begin()

		currTransaction := new(entity.Transaction)
		s.TransactionRepository.Repository.FindById(tx, currTransaction, transaction.ID)
		if currTransaction.FailedAt == nil && currTransaction.Status != 3 {
			currTransaction.Status = 3
			timeNow := time.Now()
			currTransaction.SucceedAt = &timeNow
			s.TransactionRepository.Repository.Update(tx, currTransaction)

			senderAccount := new(entity.Account)
			s.AccountRepository.Repository.FindById(tx, senderAccount, currTransaction.SenderAccountID)
			senderAccount.BalanceIDR -= currTransaction.IDRAmount
			s.AccountRepository.Repository.Update(tx, senderAccount)

			recipientAccount := new(entity.Account)
			s.AccountRepository.Repository.FindById(tx, recipientAccount, currTransaction.RecipientAccountID)
			recipientAccount.BalanceIDR += currTransaction.IDRAmount
			s.AccountRepository.Repository.Update(tx, recipientAccount)

			tx.Commit()
		}
	})

	response := &model.TransferResponse{
		TransactionID:          transaction.ID,
		RecipientAccountNumber: recipientAccount.Number,
		RecipientAccountName:   fmt.Sprintf("%s %s", recipientUser.FirstName, recipientUser.LastName),
		SenderAccountNumber:    senderAccount.Number,
		SenderAccountName:      fmt.Sprintf("%s %s", senderUser.FirstName, senderUser.LastName),
		TransferAt:             transaction.CreatedAt,
		Status:                 "pending",
		CurrencyCode:           currency.Code,
		Amount:                 transaction.OriginalAmount,
		AmountInIDR:            transaction.IDRAmount,
	}
	return response, nil
}

func (s *TransactionServiceImpl) Withdraw(ctx context.Context, request *model.WithDrawRequest) (*model.WithdrawResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("[%d] invalid request body : %+v", http.StatusBadRequest, err)
		return nil, err
	}

	tx := s.DB.WithContext(ctx).Begin()

	transaction := new(entity.Transaction)
	if err := s.TransactionRepository.Repository.FindById(tx, transaction, request.TransactionID); err != nil {
		s.Log.Warnf("[%d] transaction not found : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find transaction")
	}

	senderAccount := new(entity.Account)
	senderAccount.ID = transaction.SenderAccountID
	err := s.AccountRepository.FindByIDWithAccountType(tx, senderAccount)
	if err != nil {
		s.Log.Warnf("[%d] sender account not found : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	if request.UserID != senderAccount.UserID {
		s.Log.Warnf("[%d] invalid sender account : %+v", http.StatusNotFound, err)
		return nil, exception.NewHttpError(http.StatusNotFound, "can't find sender account")
	}

	response := new(model.WithdrawResponse)
	// if transaction is made before current time + 30 second, and if transaction still pending
	// and transaction hasn't been succeed yet, withdraw this transaction by setting status to 3 (FAILED/WITHDRAW)

	fmt.Println("TIMENOW", time.Now().After(transaction.CreatedAt.Add(time.Second*30)))
	fmt.Println(transaction.Status)
	fmt.Println(transaction.SucceedAt)

	if time.Now().Before(transaction.CreatedAt.Add(time.Second*30)) &&
		transaction.Status == 1 && transaction.SucceedAt == nil {
		transaction.Status = 3
		timeNow := time.Now()
		transaction.FailedAt = &timeNow

		if err := s.TransactionRepository.Repository.Update(tx, transaction); err != nil {
			if errRollback := tx.Rollback().Error; errRollback != nil {
				s.Log.Warnf("[%d] %+v", http.StatusInternalServerError, errRollback)
				return nil, exception.NewHttpError(http.StatusInternalServerError, "something wrong")
			}

			s.Log.Warnf("[%d] %+v", http.StatusInternalServerError, err)
			return nil, exception.NewHttpError(http.StatusInternalServerError, "something wrong")
		}

		response = &model.WithdrawResponse{
			TransactionID:  transaction.ID,
			WithdrawStatus: "success",
			WithdrawAt:     transaction.FailedAt,
			TransferAt:     transaction.CreatedAt,
			TransferStatus: "withdrawed",
		}
	} else {
		response = &model.WithdrawResponse{
			TransactionID:  transaction.ID,
			WithdrawStatus: "failed",
			TransferAt:     transaction.CreatedAt,
			TransferStatus: "succeed",
		}
	}

	return response, nil
}
