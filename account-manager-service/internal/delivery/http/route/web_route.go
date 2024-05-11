package route

import (
	"account-manager-service/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App             *gin.Engine
	ControllerSetup *controller.ControllerSetup
}

func (c *RouteConfig) Setup() {
	route := c.App.Group("/api/v1")
	c.SetupAuthRoute(route)
}

func (c *RouteConfig) SetupAuthRoute(route *gin.RouterGroup) {
	authRoute := route.Group("/auth")
	authRoute.POST("/sign-up", c.ControllerSetup.AuthController.SignUp)
	authRoute.POST("/sign-in", c.ControllerSetup.AuthController.SignIn)
}
