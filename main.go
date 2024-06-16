package main

import (
	"belio-api/config"
	"belio-api/controllers"
	"belio-api/repositories"
	"belio-api/routes"
	"belio-api/services"
	"belio-api/utils"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	config.InitDB()
	utils.MigrateDB()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	linkRepository := repositories.NewLinkRepository(config.DB)
	linkService := services.NewLinkService(linkRepository)
	linkController := controllers.NewLinkController(linkService)

	// Setup router
	router := routes.SetupRouter(userController, linkController)

	router.Run(":8080")

}
