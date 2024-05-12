package service

import (
	"account-manager-service/internal/model"
	"context"
)

type TransactionService interface {
	FindTransactions(ctx context.Context, request *model.GetUserTransactionsRequest) ([]model.TransactionResponse, error)
}
