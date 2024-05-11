package config

import (
	"account-manager-service/database/migration"
	"account-manager-service/database/seeder"
	"account-manager-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigBootstrap struct {
	ViperConfig *viper.Viper
	Logger      *logrus.Logger
	DB          *gorm.DB
	App         *gin.Engine
	Validate    *validator.Validate
}

func Bootstrap(cfg *ConfigBootstrap) {
	repositorySetup := repository.Setup()

	// Setup database
	if err := migration.Drop(cfg.DB); err != nil {
		cfg.Logger.Fatalf("failed to drop the database : %+v", err)
	}
	if err := migration.Migrate(cfg.DB); err != nil {
		cfg.Logger.Fatalf("failed to migrate the database : %+v", err)
	}
	if err := seeder.Seed(cfg.DB, repositorySetup); err != nil {
		cfg.Logger.Fatalf("failed to seed the database : %+v", err)
	}
}
