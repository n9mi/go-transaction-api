package entity

import "time"

type AccountType struct {
	ID string `gorm:"primaryKey"`

	Name     string
	LimitIDR float64 `gorm:"column:limit_idr"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
