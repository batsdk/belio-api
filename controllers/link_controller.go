package controllers

import (
	"belio-api/models"
	"belio-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LinkController struct {
	service *services.LinkService
}

func NewLinkController(service *services.LinkService) *LinkController {
	return &LinkController{
		service: service,
	}
}

func (lc *LinkController) CreateLink(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var link models.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link.UserID = uint(userId)
	if err := lc.service.Create(&link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"link": link})
}

func (lc *LinkController) GetLinkById(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	links, err := lc.service.FindLinkByUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, links)
}
