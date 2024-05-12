package repository

import (
	"payment-manager-service/internal/entity"

	"gorm.io/gorm"
)

type AccountRepository struct {
	Repository[entity.Account]
}

func NewAccountRepository() *AccountRepository {
	return new(AccountRepository)
}

func (r *AccountRepository) FindByIDWithAccountType(tx *gorm.DB, account *entity.Account) error {
	return tx.Where("accounts.id = ?", account.ID).Preload("AccountType").First(account).Error
}

func (r *AccountRepository) FindByAccountNumber(tx *gorm.DB, account *entity.Account) error {
	return tx.First(account, "number = ?", account.Number).Error
}
