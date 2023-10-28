package controllers

import (
	"gin-photo-api/app"
	"gin-photo-api/database"
	"gin-photo-api/helper"
	"gin-photo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	payload := app.UserRegisterInput{}

	c.BindJSON(&payload)

	// validate incoming body request
	err := app.ValidateUserRegister(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	// hash password
	hashedPassword, _ := helper.HashPassword(payload.Password)

	newUser := models.User{
		UserSchema: models.UserSchema{
			Username: payload.Username,
			Email:    payload.Email,
			Password: hashedPassword,
		},
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "server error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success", "data": newUser})
}

func UserList(c *gin.Context) {
	var users []models.User

	database.DB.Select("Username", "Email").Find(&users)

	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}

func UserRetrive(c *gin.Context) {
	id := c.Param("userId")
	var user models.User
	result := database.DB.Where("id = ?", id).Preload("Photo").First(&user)

	// error handling
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		// Record not found
	// 		// Handle this case
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
	// 	} else {
	// 		// Other error occurred
	// 		// Handle the error
	// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	// 	}
	// 	return
	// }
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func UserUpdate(c *gin.Context) {
	id := c.Param("userId")
	var user models.User
	payload := app.UserUpdateInput{}

	c.ShouldBindJSON(&payload)

	// validate incoming body request
	err := app.ValidateUserUpdate(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	// check if user exist
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	// hash password
	hashedPassword, _ := helper.HashPassword(payload.Password)

	// update record
	newUser := models.User{
		UserSchema: models.UserSchema{
			Username: payload.Username,
			Email:    payload.Email,
			Password: hashedPassword,
		},
	}

	// save to DB
	result = database.DB.Model(&user).Updates(newUser)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "data": newUser})
}

func UserDelete(c *gin.Context) {
	id := c.Param("userId")

	result := database.DB.Delete(&models.User{}, id)
	// error handling
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Record Deleted"})
}
