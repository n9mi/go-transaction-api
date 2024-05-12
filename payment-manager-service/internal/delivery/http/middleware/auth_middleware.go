package middleware

import (
	"net/http"
	"payment-manager-service/internal/delivery/http/exception"
	"payment-manager-service/internal/model"
	"payment-manager-service/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AuthMiddleware(viperCfg *viper.Viper, log *logrus.Logger, restyClient *resty.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := util.ExtractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		authData := new(model.AuthData)
		result, err := restyClient.R().
			SetHeader("Content-Type", "application/json").
			SetAuthToken(accessToken).
			SetResult(authData).
			Post(viperCfg.GetString("AUTH_SERVICE_URL"))

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		if !result.IsSuccess() {
			err := exception.NewHttpError(http.StatusUnauthorized, "unauthorized")
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("authData", *authData)
		c.Next()
	}
}
