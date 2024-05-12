package controller

import (
	"account-manager-service/internal/delivery/http/exception"
	"account-manager-service/internal/model"
	"account-manager-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ct *AuthController) SignUp(c *gin.Context) {
	var request model.SignUpRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(exception.NewHttpError(http.StatusBadRequest, "invalid request"))
		c.Abort()
		return
	}

	if err := ct.AuthService.SignUp(c.Request.Context(), &request); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

func (ct *AuthController) SignIn(c *gin.Context) {
	var request model.SignInRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(exception.NewHttpError(http.StatusBadRequest, "invalid request"))
		c.Abort()
		return
	}

	response, err := ct.AuthService.SignIn(c.Request.Context(), &request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ct *AuthController) RetrieveAuthData(c *gin.Context) {
	authData, ok := c.MustGet("authData").(model.AuthData)
	if !ok {
		c.Error(exception.NewHttpError(http.StatusForbidden, "invalid authorization data"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, authData)
}
