package controller

import "account-manager-service/internal/service"

type ControllerSetup struct {
	AuthController        *AuthController
	AccountController     *AccountController
	TransactionController *TransactionController
}

func Setup(serviceSetup *service.ServiceSetup) *ControllerSetup {
	return &ControllerSetup{
		AuthController:        NewAuthController(serviceSetup.AuthService),
		AccountController:     NewAccountController(serviceSetup.AccountService),
		TransactionController: NewTransactionController(serviceSetup.TransactionService),
	}
}
