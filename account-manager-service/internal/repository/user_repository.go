package repository

import (
	"account-manager-service/internal/entity"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (r *UserRepository) ExistByEmail(tx *gorm.DB, email string) (bool, error) {
	var exist bool
	err := tx.Model(new(entity.User)).Select("count(email) > 0").
		Where("email = ?", strings.ToLower(email)).
		Find(&exist).Error

	return exist, err
}

func (r *UserRepository) ExistByPhoneNumber(tx *gorm.DB, phoneNumber string) (bool, error) {
	var exist bool
	err := tx.Model(new(entity.User)).Select("count(phone_number) > 0").
		Where("phone_number = ?", strings.ToLower(phoneNumber)).
		Find(&exist).Error

	return exist, err
}

func (r *UserRepository) FindByEmail(tx *gorm.DB, user *entity.User) error {
	return tx.First(user, "email = ?", strings.ToLower(user.Email)).Error
}
