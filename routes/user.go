package routes

import (
	"belio-api/controllers"
	"github.com/gin-gonic/gin"
)

func LinkRoutes(router *gin.RouterGroup, linkController *controllers.LinkController) {
	router.POST("/users/:userId/link", linkController.CreateLink)
	router.GET("/users/:userId/link", linkController.GetLinkById)
}
