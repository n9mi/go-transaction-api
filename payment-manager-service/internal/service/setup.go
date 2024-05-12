package service

import (
	"payment-manager-service/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ServiceSetup struct {
	TransactionService TransactionService
}

func Setup(viperCfg *viper.Viper, validate *validator.Validate, db *gorm.DB, log *logrus.Logger,
	repositorySetup *repository.RepositorySetup) *ServiceSetup {
	return &ServiceSetup{
		TransactionService: NewTransactionService(db, validate, log, repositorySetup.UserRepository,
			repositorySetup.TransactionRepository, repositorySetup.AccountRepository, repositorySetup.CurrencyRepository),
	}
}
