package middleware

import (
	"account-manager-service/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AuthMiddleware(viperCfg *viper.Viper, redisClient *redis.Client, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := util.ExtractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		authData, err := util.VerifyAccessToken(c.Request.Context(), viperCfg, redisClient, log, accessToken)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("authData", *authData)
		c.Next()
	}
}
