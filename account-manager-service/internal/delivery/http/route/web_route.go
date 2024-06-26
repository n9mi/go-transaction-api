package route

import (
	"account-manager-service/internal/delivery/http/controller"
	"account-manager-service/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App             *gin.Engine
	MiddlewareSetup *middleware.MiddlewareSetup
	ControllerSetup *controller.ControllerSetup
}

func (c *RouteConfig) Setup() {
	route := c.App.Group("/api/v1")
	c.SetupAuthRoute(route)
	c.SetupProtectedAuthRoute(route)
	c.SetupAccountRoute(route)
	c.SetupTransactionRoute(route)
}

func (c *RouteConfig) SetupAuthRoute(route *gin.RouterGroup) {
	authRoute := route.Group("/auth")
	authRoute.POST("/sign-up", c.ControllerSetup.AuthController.SignUp)
	authRoute.POST("/sign-in", c.ControllerSetup.AuthController.SignIn)
}

func (c *RouteConfig) SetupProtectedAuthRoute(route *gin.RouterGroup) {
	accountRoute := route.Group("/auth/protected")
	accountRoute.Use(c.MiddlewareSetup.AuthMiddleware)
	accountRoute.POST("/verify", c.ControllerSetup.AuthController.RetrieveAuthData)
}

func (c *RouteConfig) SetupAccountRoute(route *gin.RouterGroup) {
	accountRoute := route.Group("/accounts")
	accountRoute.Use(c.MiddlewareSetup.AuthMiddleware)
	accountRoute.GET("", c.ControllerSetup.AccountController.GetAccounts)
}

func (c *RouteConfig) SetupTransactionRoute(route *gin.RouterGroup) {
	accountRoute := route.Group("/transactions")
	accountRoute.Use(c.MiddlewareSetup.AuthMiddleware)
	accountRoute.GET("", c.ControllerSetup.TransactionController.GetTransactions)
}
