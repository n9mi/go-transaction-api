package controller

import "account-manager-service/internal/service"

type ControllerSetup struct {
	AuthController *AuthController
}

func Setup(serviceSetup *service.ServiceSetup) *ControllerSetup {
	return &ControllerSetup{
		AuthController: NewAuthController(serviceSetup.AuthService),
	}
}
