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

func (r *UserRepository) FindUserById(id uint) error {
	var user models.User
	err := r.DB.First(&user, id).Error

	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepository) UpdateUserProfileImage(id uint, link string) (models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	user.ProfilePhoto = link
	if err := r.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
