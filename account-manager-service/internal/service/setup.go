package service

import (
	"account-manager-service/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ServiceSetup struct {
	AuthService        AuthService
	AccountService     AccountService
	TransactionService TransactionService
}

func Setup(viperCfg *viper.Viper, validate *validator.Validate, db *gorm.DB, redisClient *redis.Client, log *logrus.Logger,
	repositorySetup *repository.RepositorySetup) *ServiceSetup {
	return &ServiceSetup{
		AuthService:        NewAuthService(viperCfg, validate, db, redisClient, log, repositorySetup.UserRepository),
		AccountService:     NewAccountService(db, log, repositorySetup.UserRepository, repositorySetup.AccountRepository),
		TransactionService: NewTransactionService(db, log, repositorySetup.UserRepository, repositorySetup.TransactionRepository),
	}
}
