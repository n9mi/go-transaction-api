package controller

import "account-manager-service/internal/service"

type ControllerSetup struct {
	AuthController    *AuthController
	AccountController *AccountController
}

func Setup(serviceSetup *service.ServiceSetup) *ControllerSetup {
	return &ControllerSetup{
		AuthController:    NewAuthController(serviceSetup.AuthService),
		AccountController: NewAccountController(serviceSetup.AccountService),
	}
}
