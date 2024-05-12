package model

import "time"

type GetUserAccountsRequest struct {
	Page          int
	PageSize      int
	UserID        string
	AccountTypeID string
}

type AccountResponse struct {
	AccountTypeID     string    `json:"account_type_id"`
	AccountTypeName   string    `json:"account_type_name"`
	AccountNumber     string    `json:"account_number"`
	AccountBalanceIDR float64   `json:"account_balance_idr"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
