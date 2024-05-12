package service

import (
	"account-manager-service/internal/model"
	"account-manager-service/internal/repository"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountServiceImpl struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	UserRepository        *repository.UserRepository
	AccountRepository     *repository.AccountRepository
	TransactionRepository *repository.TransactionRepository
}

func NewAccountService(db *gorm.DB, log *logrus.Logger, userRepository *repository.UserRepository,
	accountRepository *repository.AccountRepository, transactionRepository *repository.TransactionRepository) AccountService {
	return &AccountServiceImpl{
		DB:                    db,
		Log:                   log,
		UserRepository:        userRepository,
		AccountRepository:     accountRepository,
		TransactionRepository: transactionRepository,
	}
}

func (s *AccountServiceImpl) FindAccounts(ctx context.Context, request *model.GetUserAccountsRequest) ([]model.AccountResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	accounts, err := s.AccountRepository.FindByUserID(tx, request)
	accountsResponse := make([]model.AccountResponse, len(accounts))
	for i, acE := range accounts {
		accountsResponse[i] = model.AccountResponse{
			AccountTypeID:     acE.AccountTypeID,
			AccountTypeName:   acE.AccountType.Name,
			AccountNumber:     acE.Number,
			AccountBalanceIDR: acE.BalanceIDR,
			CreatedAt:         acE.CreatedAt,
			UpdatedAt:         *acE.UpdatedAt,
		}
	}

	return accountsResponse, err
}
