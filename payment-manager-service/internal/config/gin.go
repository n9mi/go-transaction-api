package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewGin(viperConfig *viper.Viper) *gin.Engine {
	r := gin.Default()

	return r
}
