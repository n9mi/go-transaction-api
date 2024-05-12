package repository

import "payment-manager-service/internal/entity"

type AccountTypeRepository struct {
	Repository[entity.AccountType]
}

func NewAccounTypeRepository() *AccountTypeRepository {
	return new(AccountTypeRepository)
}
