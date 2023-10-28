package controllers

import (
	"gin-photo-api/app"
	"gin-photo-api/database"
	"gin-photo-api/helper"
	"gin-photo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PhotoCreate(c *gin.Context) {
	payload := app.PhotoCreateInput{}
	var user models.User

	c.ShouldBindJSON(&payload)

	// validate incoming body request
	err := app.ValidatePhotoCreate(payload)
	if err != nil {
		helper.ValidationError(err, c)
		return
	}

	//check the respective user
	result := database.DB.Where("email = ?", payload.UserEmail).First(&user)
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}

	//insert to DB
	newPhoto := models.Photo{
		PhotoSchema: models.PhotoSchema{
			Title:    payload.Title,
			Caption:  payload.Caption,
			PhotoURL: payload.PhotoURL,
		},
		UserID: user.ID,
	}
	result = database.DB.Create(&newPhoto)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
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

	result := database.DB.Delete(&models.Photo{}, id)
	// error handling
	if result.Error != nil {
		helper.RecordNotFoundError(result, c)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Record Deleted"})
}
