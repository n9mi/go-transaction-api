package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareSetup struct {
	AuthMiddleware gin.HandlerFunc
}

func NewMiddlewareSetup(viperCfg *viper.Viper, redisClient *redis.Client, log *logrus.Logger) *MiddlewareSetup {
	return &MiddlewareSetup{
		AuthMiddleware: AuthMiddleware(viperCfg, redisClient, log),
	}
}
