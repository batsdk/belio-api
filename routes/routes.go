package routes

import (
	"belio-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users", userController.GetUsers)
		api.GET("/test", userController.Test)
	}

	return router

}
