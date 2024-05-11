package migration

import (
	"account-manager-service/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.AccountType{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.Currency{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.Account{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.Transaction{}); err != nil {
		return err
	}

	return nil
}
