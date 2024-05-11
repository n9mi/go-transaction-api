package repository

import "account-manager-service/internal/entity"

type CurrencyRepository struct {
	Repository[entity.Currency]
}

func NewCurrencyRepository() *CurrencyRepository {
	return new(CurrencyRepository)
}
