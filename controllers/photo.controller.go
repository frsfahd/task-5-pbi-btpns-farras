package controllers

import (
	"gin-photo-api/app"
	"gin-photo-api/database"
	"gin-photo-api/helper"
	"gin-photo-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PhotoCreate(c *gin.Context) {
	payload := app.PhotoCreateInput{}

	// get current user
	user, _ := c.Get("user")
	currentUser, _ := user.(models.User)

	c.ShouldBindJSON(&payload)

	// validate incoming body request
	err := app.ValidatePhotoCreate(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	//check the respective user
	// result := database.DB.Where("email = ?", payload.UserEmail).First(&user)
	// if result.Error != nil {
	// 	helper.RecordNotFoundError(result, c)
	// 	return
	// }

	//insert to DB
	newPhoto := models.Photo{
		PhotoSchema: models.PhotoSchema{
			Title:    payload.Title,
			Caption:  payload.Caption,
			PhotoURL: payload.PhotoURL,
		},
		UserID: currentUser.ID,
	}
	result := database.DB.Create(&newPhoto)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success", "data": newPhoto})
}

func PhotoList(c *gin.Context) {
	var photos []models.Photo

	result := database.DB.Find(&photos)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "data": photos})
}

func PhotoRetrive(c *gin.Context) {
	id := c.Param("photoId")

	var photo models.Photo
	result := database.DB.Where("id = ?", id).First(&photo)

	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "data": photo})
}

func PhotoUpdate(c *gin.Context) {
	id := c.Param("photoId")
	var photo models.Photo

	payload := app.PhotoUpdateInput{}

	c.ShouldBindJSON(&payload)

	// validate incoming body request
	err := app.ValidatePhotoUpdate(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	//check the respective user
	// var user models.User
	// database.DB.Where("email = ?", payload.UserEmail).First(&user)
	// if gorm.ErrRecordNotFound != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	// }

	// check if photo exist
	result := database.DB.Where("id = ?", id).First(&photo)
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	// get current user
	id = strconv.FormatUint(uint64(photo.UserID), 10)
	_, ok := helper.ValidateCurrentUser(id, c)
	if !ok {
		return
	}

	// update record
	newPhoto := models.Photo{
		PhotoSchema: models.PhotoSchema{
			Title:    payload.Title,
			Caption:  payload.Caption,
			PhotoURL: payload.PhotoURL,
		},
	}

	result = database.DB.Model(&photo).Updates(newPhoto)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "data": photo})
}

func PhotoDelete(c *gin.Context) {
	id := c.Param("photoId")

	var photo models.Photo
	result := database.DB.Where("id = ?", id).First(&photo)
	// error handling
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	// validate auth
	id = strconv.FormatUint(uint64(photo.UserID), 10)
	_, ok := helper.ValidateCurrentUser(id, c)
	if !ok {
		return
	}

	result = database.DB.Delete(&photo)
	// error handling
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Record Deleted"})
}
