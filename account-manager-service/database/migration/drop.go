package migration

import (
	"account-manager-service/internal/entity"

	"gorm.io/gorm"
)

func Drop(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&entity.Transaction{}); err != nil {
		return err
	}

	if err := db.Migrator().DropTable(&entity.Account{}); err != nil {
		return err
	}

	if err := db.Migrator().DropTable(&entity.Currency{}); err != nil {
		return err
	}

	if err := db.Migrator().DropTable(&entity.AccountType{}); err != nil {
		return err
	}

	if err := db.Migrator().DropTable(&entity.User{}); err != nil {
		return err
	}

	return nil
}
