package repository

import (
	"gin-login/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *entity.User, email string) error
	CreateUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *entity.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *entity.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
