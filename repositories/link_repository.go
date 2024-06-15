package repositories

import (
	"belio-api/models"

	"gorm.io/gorm"
)

type LinkRepositories interface {
	Create(link *models.Link) error
	FindLinkByUser(subID string) ([]models.Link, error)
}

type LinkRepository struct {
	DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		DB: db,
	}
}

func (r *LinkRepository) Create(link *models.Link) error {
	return r.DB.Create(link).Error
}

func (r *LinkRepository) FindLinkByUser(subID uint) ([]models.Link, error) {
	var links []models.Link
	err := r.DB.Where("user_id = ?", subID).Find(&links).Error
	return links, err
}
