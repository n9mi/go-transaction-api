package controller

import (
	"net/http"
	"payment-manager-service/internal/delivery/http/exception"
	"payment-manager-service/internal/model"
	"payment-manager-service/internal/service"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: transactionService,
	}
}

func (ct *TransactionController) Transfer(c *gin.Context) {
	authData, ok := c.MustGet("authData").(model.AuthData)
	if !ok {
		c.Error(exception.NewHttpError(http.StatusForbidden, "invalid authorization data"))
		c.Abort()
		return
	}

	var request model.TransferRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(exception.NewHttpError(http.StatusBadRequest, "invalid request"))
		c.Abort()
		return
	}

	request.UserID = authData.UserID
	data, err := ct.TransactionService.Transfer(c.Request.Context(), &request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ct *TransactionController) Withdraw(c *gin.Context) {
	authData, ok := c.MustGet("authData").(model.AuthData)
	if !ok {
		c.Error(exception.NewHttpError(http.StatusForbidden, "invalid authorization data"))
		c.Abort()
		return
	}

	request := &model.WithDrawRequest{
		TransactionID: c.Param("transactionId"),
		UserID:        authData.UserID,
	}
	data, err := ct.TransactionService.Withdraw(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
