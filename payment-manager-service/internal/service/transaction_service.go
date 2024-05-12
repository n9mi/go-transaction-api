package service

import (
	"context"
	"payment-manager-service/internal/model"
)

type TransactionService interface {
	Transfer(ctx context.Context, request *model.TransferRequest) (*model.TransferResponse, error)
}
