package repository

import "account-manager-service/internal/entity"

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}
