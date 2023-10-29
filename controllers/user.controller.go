package controllers

import (
	"errors"
	"gin-photo-api/app"
	"gin-photo-api/database"
	"gin-photo-api/helper"
	"gin-photo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// error handling
	if result.Error != nil {
		helper.DuplicateUserError(result, c)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success", "data": newUser})
}

func UserLogin(c *gin.Context) {
	// get email and password
	payload := app.UserLoginInput{}

	c.BindJSON(&payload)

	// validate incoming body request
	err := app.ValidateUserLogin(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	// look up requested user
	var user models.User
	result := database.DB.Where("email = ?", payload.Email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Email"})
		}
		return
	}

	// compare sent in password with actual password
	ok := helper.CheckPasswordHash(payload.Password, user.Password)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Password"})
		return
	}

	// generate token
	token := helper.GenerateToken(payload.Email)

	// send token back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user.ID})
}

func UserList(c *gin.Context) {
	var users []models.User

	database.DB.Omit("Photo").Find(&users)

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

	// get current user
	currentUser, ok := helper.ValidateCurrentUser(id, c)
	if !ok {
		return
	}

	payload := app.UserUpdateInput{}

	c.ShouldBindJSON(&payload)

	// validate incoming body request
	err := app.ValidateUserUpdate(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	// hash password if password field not empty
	password := payload.Password
	if len(password) != 0 {
		password, _ = helper.HashPassword(payload.Password)
	}

	// update record
	newUser := models.User{
		UserSchema: models.UserSchema{
			Username: payload.Username,
			Email:    payload.Email,
			Password: password,
		},
	}

	// save to DB
	result := database.DB.Model(&currentUser).Updates(newUser)

	// error handling
	if result.Error != nil {
		helper.DuplicateUserError(result, c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "data": newUser})
}

func UserDelete(c *gin.Context) {
	id := c.Param("userId")

	// get current user
	currentUser, ok := helper.ValidateCurrentUser(id, c)
	if !ok {
		return
	}

	result := database.DB.Delete(&currentUser)

	// error handling
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Record Deleted"})
}
