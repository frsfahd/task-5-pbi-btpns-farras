package routes

import (
	"gin-photo-api/controllers"
	"gin-photo-api/middleware"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")

	// user.GET("/", controllers.UserList) // testing
	// user.GET("/:userId", controllers.UserRetrive) // testing
	user.POST("/register", controllers.UserCreate)
	user.POST("/login", controllers.UserLogin)
	user.PUT("/:userId", middleware.Auth, controllers.UserUpdate)    //protected
	user.DELETE("/:userId", middleware.Auth, controllers.UserDelete) //protected

}
