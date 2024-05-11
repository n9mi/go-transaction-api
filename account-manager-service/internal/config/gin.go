package config

import (
	"account-manager-service/internal/delivery/http/exception"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewGin(viperCfg *viper.Viper) *gin.Engine {
	r := gin.Default()
	r.Use(customErrorHandler())

	return r
}

func customErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			var httpCode int
			var errorMessages []string

			switch e := err.Err.(type) {
			case validator.ValidationErrors:
				httpCode = http.StatusBadRequest
				for _, errItem := range e {
					switch errItem.Tag() {
					case "required":
						errorMessages = append(errorMessages,
							fmt.Sprintf("%s is required", errItem.Field()))
					case "min":
						errorMessages = append(errorMessages,
							fmt.Sprintf("%s is should be more than %s character", errItem.Field(), errItem.Param()))
					case "max":
						errorMessages = append(errorMessages,
							fmt.Sprintf("%s is should be less than %s character", errItem.Field(), errItem.Param()))
					case "email":
						errorMessages = append(errorMessages,
							fmt.Sprintf("%s should be a valid email", errItem.Field()))
					}
				}
			case *exception.HttpError:
				httpCode = e.Code
				errorMessages = e.Messages
			default:
				httpCode = http.StatusInternalServerError
				errorMessages = append(errorMessages, e.Error())
			}

			fmt.Print(httpCode, errorMessages)

			c.JSON(httpCode, gin.H{
				"messages": errorMessages,
			})
			c.Abort()
		}
	}
}
