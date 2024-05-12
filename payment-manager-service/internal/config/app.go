package config

import (
	"payment-manager-service/internal/delivery/http/controller"
	"payment-manager-service/internal/delivery/http/middleware"
	"payment-manager-service/internal/delivery/http/route"
	"payment-manager-service/internal/repository"
	"payment-manager-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
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
	RestyClient *resty.Client
}

func Bootstrap(cfg *ConfigBootstrap) {
	repositorySetup := repository.Setup()

	serviceSetup := service.Setup(
		cfg.ViperConfig,
		cfg.Validate,
		cfg.DB,
		cfg.Log,
		repositorySetup)

	controllerSetup := controller.Setup(
		serviceSetup,
	)

	middlewareSetup := middleware.NewMiddlewareSetup(
		cfg.ViperConfig,
		cfg.Log,
		cfg.RestyClient,
	)

	routeConfig := route.RouteConfig{
		App:             cfg.App,
		MiddlewareSetup: middlewareSetup,
		ControllerSetup: controllerSetup,
	}
	routeConfig.Setup()
}
