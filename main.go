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

	router := routes.SetupRouter(userController)
	router.Run(":8080")

}
