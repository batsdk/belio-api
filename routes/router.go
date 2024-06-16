package routes

import (
	"belio-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, linkController *controllers.LinkController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		UserRoutes(api, userController)
		LinkRoutes(api, linkController)
	}

	return router
}
