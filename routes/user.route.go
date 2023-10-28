package routes

import (
	"gin-photo-api/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")

	user.GET("/", controllers.UserList)
	user.GET("/:userId", controllers.UserRetrive)
	user.POST("/register", controllers.UserCreate)
	user.PUT("/:userId", controllers.UserUpdate)
	user.DELETE("/:userId", controllers.UserDelete)
}
