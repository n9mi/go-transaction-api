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
	AuthService AuthService
}

func Setup(viperCfg *viper.Viper, validate *validator.Validate, db *gorm.DB, redisClient *redis.Client, log *logrus.Logger,
	repositorySetup *repository.RepositorySetup) *ServiceSetup {
	return &ServiceSetup{
		AuthService: NewAuthService(viperCfg, validate, db, redisClient, log, repositorySetup.UserRepository),
	}
}
