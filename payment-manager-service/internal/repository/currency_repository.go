package repository

import (
	"payment-manager-service/internal/entity"

	"gorm.io/gorm"
)

type CurrencyRepository struct {
	Repository[entity.Currency]
}

func NewCurrencyRepository() *CurrencyRepository {
	return new(CurrencyRepository)
}

func (r *CurrencyRepository) FindCurrentCurrencyByCode(tx *gorm.DB, currency *entity.Currency) error {
	return tx.Order("updated_at desc").First(currency, "code = ?", currency.Code).Error
}
