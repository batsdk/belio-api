package controllers

import (
	"belio-api/models"
	"belio-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})

}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (ctrl *UserController) UploadProfilePhoto(c *gin.Context) {

	//	Find if the user with ID exists
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	//	Get the Photo from req

	//	Save the file locally

	//	Save to CDN

	//	Update User's profile Photo
	user, err := ctrl.service.UpdateUserProfileImage(uint(userId), "somelinkhere")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func (ctrl *UserController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tests": "test update ar 8.35"})
}
