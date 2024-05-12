package controller

import "payment-manager-service/internal/service"

type ControllerSetup struct {
	TransactionController *TransactionController
}

func Setup(serviceSetup *service.ServiceSetup) *ControllerSetup {
	return &ControllerSetup{
		TransactionController: NewTransactionController(serviceSetup.TransactionService),
	}
}
