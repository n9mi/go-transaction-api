package repository

import (
	"account-manager-service/internal/entity"
	"account-manager-service/internal/model"
	"account-manager-service/internal/util"

	"gorm.io/gorm"
)

type AccountRepository struct {
	Repository[entity.Account]
}

func NewAccountRepository() *AccountRepository {
	return new(AccountRepository)
}

func (r *AccountRepository) FindByUserID(tx *gorm.DB, request *model.GetUserAccountsRequest) ([]entity.Account, error) {
	var accounts []entity.Account

	if request.Page > 0 && request.PageSize > 0 {
		tx = tx.Scopes(util.Paginate(request.Page, request.PageSize))
	}

	tx = tx.Select("number", "balance_idr", "accounts.account_type_id", "accounts.created_at", "accounts.updated_at")
	tx = tx.Joins("AccountType").Where("user_id = ?", request.UserID).Find(&accounts)

	if len(request.AccountTypeID) > 0 {
		tx = tx.Where("accounts.account_type_id = ?", request.AccountTypeID)
	}

	return accounts, tx.Error
}
