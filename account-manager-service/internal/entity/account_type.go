package entity

import "time"

type AccountType struct {
	ID string `gorm:"primaryKey"`

	Name string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
