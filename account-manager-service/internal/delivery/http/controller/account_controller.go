package controller

import (
	"account-manager-service/internal/delivery/http/exception"
	"account-manager-service/internal/model"
	"account-manager-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) *AccountController {
	return &AccountController{
		AccountService: accountService,
	}
}

func (ct *AccountController) GetAccounts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	authData, ok := c.MustGet("authData").(model.AuthData)
	if !ok {
		c.Error(exception.NewHttpError(http.StatusForbidden, "invalid authorization data"))
		c.Abort()
		return
	}
	request := &model.GetUserAccountsRequest{
		Page:          page,
		PageSize:      pageSize,
		UserID:        authData.UserID,
		AccountTypeID: c.DefaultQuery("accountTypeId", ""),
	}

	data, err := ct.AccountService.FindAccounts(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
