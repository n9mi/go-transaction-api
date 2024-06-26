package service

import (
	"account-manager-service/internal/model"
	"context"
)

type AuthService interface {
	SignUp(ctx context.Context, request *model.SignUpRequest) error
	SignIn(ctx context.Context, request *model.SignInRequest) (*model.TokenResponse, error)
}
