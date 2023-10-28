package routes

import (
	"gin-photo-api/controllers"

	"github.com/gin-gonic/gin"
)

func addPhotoRoutes(rg *gin.RouterGroup) {
	photo := rg.Group("/photos")

	photo.POST("/", controllers.PhotoCreate)
	photo.GET("/", controllers.PhotoList)
	photo.GET("/:photoId", controllers.PhotoRetrive)
	photo.PUT("/:photoId", controllers.PhotoUpdate)
	photo.DELETE("/:photoId", controllers.PhotoDelete)
}
