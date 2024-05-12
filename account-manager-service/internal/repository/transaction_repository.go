package repository

import (
	"account-manager-service/internal/entity"
	"account-manager-service/internal/model"
	"account-manager-service/internal/util"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
}

func NewTransactionRepository() *TransactionRepository {
	return new(TransactionRepository)
}

func (r *TransactionRepository) FindByUserID(tx *gorm.DB, request *model.GetUserTransactionsRequest) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if request.Page > 0 && request.PageSize > 0 {
		tx = tx.Scopes(util.Paginate(request.Page, request.PageSize))
	}

	if len(request.SenderAccountNumber) > 0 {
		tx = tx.Preload("SenderAccount", "number = ?", request.SenderAccountNumber)
	} else {
		tx = tx.Preload("SenderAccount")
	}

	if len(request.RecipientAccountNumber) > 0 {
		tx = tx.Preload("RecipientAccount", "number = ?", request.RecipientAccountNumber)
	} else {
		tx = tx.Preload("RecipientAccount")
	}

	tx = tx.Preload("Currency").Joins("inner join accounts on accounts.id = transactions.sender_account_id").
		Where("accounts.user_id = ?", request.UserID)

	if request.Status > 0 {
		tx = tx.Where("transactions.status = ?", request.Status)
	}
	tx = tx.Find(&transactions)

	return transactions, tx.Error
}
