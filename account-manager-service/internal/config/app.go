package config

import (
	"account-manager-service/database/migration"
	"account-manager-service/database/seeder"
	"account-manager-service/internal/delivery/http/controller"
	"account-manager-service/internal/delivery/http/route"
	"account-manager-service/internal/repository"
	"account-manager-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigBootstrap struct {
	ViperConfig *viper.Viper
	Log         *logrus.Logger
	DB          *gorm.DB
	App         *gin.Engine
	Validate    *validator.Validate
	RedisClient *redis.Client
}

func Bootstrap(cfg *ConfigBootstrap) {
	repositorySetup := repository.Setup()

	serviceSetup := service.Setup(
		cfg.ViperConfig,
		cfg.Validate,
		cfg.DB,
		cfg.RedisClient,
		cfg.Log,
		repositorySetup)

	controllerSetup := controller.Setup(
		serviceSetup,
	)

	routeConfig := route.RouteConfig{
		App:             cfg.App,
		ControllerSetup: controllerSetup,
	}
	routeConfig.Setup()

	// Setup database
	if err := migration.Drop(cfg.DB); err != nil {
		cfg.Log.Fatalf("failed to drop the database : %+v", err)
	}
	if err := migration.Migrate(cfg.DB); err != nil {
		cfg.Log.Fatalf("failed to migrate the database : %+v", err)
	}
	if err := seeder.Seed(cfg.DB, repositorySetup); err != nil {
		cfg.Log.Fatalf("failed to seed the database : %+v", err)
	}
}
