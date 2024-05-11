package config

import (
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

}
