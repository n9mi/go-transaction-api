package entity

import "time"

type Currency struct {
	ID string `gorm:"primaryKey"`

	Code           string
	CurrentPer1IDR float64

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (*Currency) TableName() string {
	return "currencies"
}
