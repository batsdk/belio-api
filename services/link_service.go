package services

import (
	"belio-api/models"
	"belio-api/repositories"
)

//type LinkService interface {
//	Create(link *models.Link) error
//	FindLinkByUser(subID string) ([]models.Link, error)
//}

type LinkService struct {
	repo *repositories.LinkRepository
}

func NewLinkService(linkRepo *repositories.LinkRepository) *LinkService {
	return &LinkService{repo: linkRepo}
}

func (service *LinkService) Create(link *models.Link) error {
	return service.repo.Create(link)
}

func (service *LinkService) FindLinkByUser(subId uint) ([]models.Link, error) {
	return service.repo.FindLinkByUser(subId)
}
