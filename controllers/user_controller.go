package controllers

import (
	"belio-api/models"
	"belio-api/services"
	"belio-api/utils"
	"net/http"
	"os"
	filepath2 "path/filepath"
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

	err = ctrl.service.FindUserById(uint(userId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with ID does not exist"})
		return
	}

	//	Get the Photo from req
	file, err := c.FormFile("profile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	//	Save the file locally
	filepath := filepath2.Join("/tmp", file.Filename)

	// Ensure the temporary file is deleted after processing
	defer os.Remove(filepath)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "detail": "Unable to save the file"})
		return
	}

	//	Save to CDN
	var cloudinaryUrl string
	cloudinaryUrl, err = utils.UploadToCloudinary(filepath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to Cloudinary"})
		return
	}

	//	Update User's profile Photo
	user, err := ctrl.service.UpdateUserProfileImage(uint(userId), cloudinaryUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func (ctrl *UserController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tests": "test update ar 8.35"})
}
