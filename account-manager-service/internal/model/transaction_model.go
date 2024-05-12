package model

import "time"

type GetUserTransactionsRequest struct {
	Page                   int
	PageSize               int
	UserID                 string
	SenderAccountNumber    string
	RecipientAccountNumber string
	Status                 int8
}

type TransactionResponse struct {
	ID                     string     `json:"id"`
	SenderAccountNumber    string     `json:"sender_account_number"`
	RecipientAccountNumber string     `json:"recipient_account_number"`
	RecipientAccountName   string     `json:"recipient_account_name"`
	CurrencyCode           string     `json:"currency_code"`
	OriginalAmount         float64    `json:"orgininal_amount"`
	AmountInIDR            float64    `json:"amount_in_idr"`
	Status                 string     `json:"status"`
	CreatedAt              time.Time  `json:"created_at"`
	SucceedAt              *time.Time `json:"succeed_at"`
	FailedAt               *time.Time `json:"failed_at"`
}
