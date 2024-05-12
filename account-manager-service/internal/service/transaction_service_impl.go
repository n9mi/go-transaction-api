package service

import (
	"account-manager-service/internal/entity"
	"account-manager-service/internal/model"
	"account-manager-service/internal/repository"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	UserRepository        *repository.UserRepository
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionService(db *gorm.DB, log *logrus.Logger, userRepository *repository.UserRepository,
	transactionRepository *repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		DB:                    db,
		Log:                   log,
		UserRepository:        userRepository,
		TransactionRepository: transactionRepository,
	}
}

func (s *TransactionServiceImpl) FindTransactions(ctx context.Context, request *model.GetUserTransactionsRequest) ([]model.TransactionResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	transactions, err := s.TransactionRepository.FindByUserID(tx, request)
	var transactionsResponse []model.TransactionResponse
	for _, tE := range transactions {
		if len(tE.RecipientAccount.Number) > 0 && len(tE.SenderAccount.Number) > 0 {
			transactionRes := model.TransactionResponse{
				ID:                     tE.ID,
				SenderAccountNumber:    tE.SenderAccount.Number,
				RecipientAccountNumber: tE.RecipientAccount.Number,
				CurrencyCode:           tE.Currency.Code,
				OriginalAmount:         tE.OriginalAmount,
				AmountInIDR:            tE.IDRAmount,
				CreatedAt:              tE.CreatedAt,
				SucceedAt:              tE.SucceedAt,
				FailedAt:               tE.FailedAt,
			}

			recipientUser := new(entity.User)
			if err := s.UserRepository.Repository.FindById(tx, recipientUser, tE.RecipientAccount.UserID); err != nil {
				transactionRes.RecipientAccountName = ""
			} else {
				transactionRes.RecipientAccountName = fmt.Sprintf("%s %s", recipientUser.FirstName, recipientUser.LastName)
			}

			switch tE.Status {
			case 1:
				transactionRes.Status = "pending"
			case 2:
				transactionRes.Status = "failed"
			case 3:
				transactionRes.Status = "success"
			}
			transactionsResponse = append(transactionsResponse, transactionRes)
		}
	}

	return transactionsResponse, err
}
