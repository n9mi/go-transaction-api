package entity

import "time"

type Transaction struct {
	ID string `gorm:"primaryKey"`

	SenderAccountID string
	SenderAccount   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RecipientAccountID string
	RecipientAccount   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CurrencyID string
	Currency   Currency `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	OriginalAmount float64
	IDRAmount      float64

	Status int8

	CreatedAt time.Time
	SucceedAt *time.Time
	FailedAt  *time.Time
}
