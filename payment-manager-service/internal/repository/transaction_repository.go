package repository

import (
	"payment-manager-service/internal/entity"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
}

func NewTransactionRepository() *TransactionRepository {
	return new(TransactionRepository)
}
