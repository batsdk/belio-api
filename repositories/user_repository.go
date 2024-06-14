package repositories

import (
	"belio-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error

	return users, err
}
