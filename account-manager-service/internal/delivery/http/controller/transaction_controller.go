package controller

import (
	"account-manager-service/internal/delivery/http/exception"
	"account-manager-service/internal/model"
	"account-manager-service/internal/service"
	"net/http"
	"strconv"
	"strings"

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

func (ct *TransactionController) GetTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	authData, ok := c.MustGet("authData").(model.AuthData)
	if !ok {
		c.Error(exception.NewHttpError(http.StatusForbidden, "invalid authorization data"))
		c.Abort()
		return
	}
	request := &model.GetUserTransactionsRequest{
		Page:                   page,
		PageSize:               pageSize,
		UserID:                 authData.UserID,
		RecipientAccountNumber: c.DefaultQuery("recipientAccountNumber", ""),
		SenderAccountNumber:    c.DefaultQuery("senderAccountNumber", ""),
	}
	switch strings.ToLower(c.DefaultQuery("status", "")) {
	case "":
		request.Status = 0
	case "pending":
		request.Status = 1
	case "failed":
		request.Status = 2
	case "success":
		request.Status = 3
	}

	data, err := ct.TransactionService.FindTransactions(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
