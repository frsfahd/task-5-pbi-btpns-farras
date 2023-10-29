package routes

import (
	"gin-photo-api/controllers"
	"gin-photo-api/middleware"

	"github.com/gin-gonic/gin"
)

func addPhotoRoutes(rg *gin.RouterGroup) {
	photo := rg.Group("/photos")

	photo.POST("/", middleware.Auth, controllers.PhotoCreate) //protected
	photo.GET("/", controllers.PhotoList)
	photo.GET("/:photoId", controllers.PhotoRetrive)
	photo.PUT("/:photoId", middleware.Auth, controllers.PhotoUpdate)    //protected
	photo.DELETE("/:photoId", middleware.Auth, controllers.PhotoDelete) //protected
}
