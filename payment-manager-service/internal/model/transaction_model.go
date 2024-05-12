package model

import "time"

type TransferRequest struct {
	UserID                 string  `json:"-"`
	SenderAccountID        string  `json:"sender_account_id" validate:"required"`
	RecipientAccountNumber string  `json:"recipient_account_number" validate:"required"`
	CurrencyCode           string  `json:"currency_code" validate:"required"`
	Amount                 float64 `json:"amount" validate:"required,gte=1"`
}

type TransferResponse struct {
	TransactionID          string    `json:"transaction_id"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
	RecipientAccountName   string    `json:"recipient_account_name"`
	SenderAccountNumber    string    `json:"sender_account_number"`
	SenderAccountName      string    `json:"sender_account_name"`
	TransferAt             time.Time `json:"transfer_at"`
	Status                 string    `json:"status"`
	CurrencyCode           string    `json:"currency_code"`
	Amount                 float64   `json:"amount"`
	AmountInIDR            float64   `json:"amount_in_idr"`
}

type WithDrawRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
}

type WithdrawResponse struct {
	TransactionID          string    `json:"transaction_id"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
	RecipientAccountName   string    `json:"recipient_account_name"`
	SenderAccountNumber    string    `json:"sender_account_number"`
	SenderAccountName      string    `json:"sender_account_name"`
	TransferAt             time.Time `json:"transfer_at"`
	WithdrawAt             time.Time `json:"withdraw_at"`
	Status                 string    `json:"status"`
	CurrencyCode           string    `json:"currency_code"`
	Amount                 float64   `json:"amount"`
	AmountInIDR            float64   `json:"amount_in_idr"`
}
