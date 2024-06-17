package services

import (
	"belio-api/models"
	"belio-api/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) FindUserById(id uint) error {
	return s.repo.FindUserById(id)
}

func (s *UserService) UpdateUserProfileImage(userId uint, cloudinaryLink string) (models.User, error) {
	return s.repo.UpdateUserProfileImage(userId, cloudinaryLink)
}
