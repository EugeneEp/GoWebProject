package repository

import "GoWebProject/internal/entity"

type UserRepository struct {
	Repository[entity.UserEntity]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
