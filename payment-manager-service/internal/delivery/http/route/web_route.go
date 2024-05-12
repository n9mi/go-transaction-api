package route

import (
	"payment-manager-service/internal/delivery/http/controller"
	"payment-manager-service/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App             *gin.Engine
	MiddlewareSetup *middleware.MiddlewareSetup
	ControllerSetup *controller.ControllerSetup
}

func (c *RouteConfig) Setup() {
	route := c.App.Group("/api/v1")
	c.SetupPaymentRoute(route)
}

func (c *RouteConfig) SetupPaymentRoute(route *gin.RouterGroup) {
	accountRoute := route.Group("/payment")
	accountRoute.Use(c.MiddlewareSetup.AuthMiddleware)
	accountRoute.POST("/transfer", c.ControllerSetup.TransactionController.Transfer)
}
