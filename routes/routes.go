package routes

import (
	"belio-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, linkController *controllers.LinkController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		//User Routes
		api.POST("/users", userController.CreateUser)
		api.GET("/users", userController.GetUsers)
		api.GET("/test", userController.Test)

		// Link Routes
		api.POST("users/:userId/link", linkController.CreateLink)
		api.GET("users/:userId/link", linkController.GetLinkById)

	}

	return router

}
