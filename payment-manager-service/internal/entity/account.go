package entity

import "time"

type Account struct {
	ID string `gorm:"primaryKey"`

	UserID string
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	AccountTypeID string
	AccountType   AccountType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Number     string
	BalanceIDR float64 `gorm:"column:balance_idr"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
