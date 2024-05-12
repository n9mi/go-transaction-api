package service

import (
	"account-manager-service/internal/model"
	"context"
)

type AccountService interface {
	FindAccounts(ctx context.Context, request *model.GetUserAccountsRequest) ([]model.AccountResponse, error)
}
