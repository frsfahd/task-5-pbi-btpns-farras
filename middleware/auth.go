package middleware

import (
	"gin-photo-api/database"
	"gin-photo-api/helper"
	"gin-photo-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {
	// get the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization"})
	}

	// validate

	token, err := helper.ParseToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expire
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Expired"})
			return
		}
		// find the corresponding user
		var user models.User
		result := database.DB.Where("email = ?", claims["sub"]).First(&user)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization"})
			return
		}
		// attach to request
		c.Set("user", user)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization"})
		return
	}

	c.Next()
}
