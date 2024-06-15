package utils

import (
	"belio-api/config"
	"belio-api/models"
)

func MigrateDB() {
	config.DB.AutoMigrate(&models.User{}, &models.Link{})
}
