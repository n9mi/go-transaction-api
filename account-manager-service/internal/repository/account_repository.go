package repository

import "account-manager-service/internal/entity"

type AccountRepository struct {
	Repository[entity.Account]
}

func NewAccountRepository() *AccountRepository {
	return new(AccountRepository)
}
