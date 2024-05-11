package entity

import (
	"time"
)

type User struct {
	ID string `gorm:"primaryKey"`

	FirstName string
	LastName  string

	Password string

	Phone string
	Email string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
