package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareSetup struct {
	AuthMiddleware gin.HandlerFunc
}

func NewMiddlewareSetup(viperCfg *viper.Viper, log *logrus.Logger, restyClient *resty.Client) *MiddlewareSetup {
	return &MiddlewareSetup{
		AuthMiddleware: AuthMiddleware(viperCfg, log, restyClient),
	}
}
