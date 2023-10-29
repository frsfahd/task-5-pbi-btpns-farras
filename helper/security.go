package helper

import (
	"fmt"
	"gin-photo-api/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(email string) string {
	key := os.Getenv("KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "gin-photo-api",
			"sub": email,
			"exp": time.Now().Add(time.Hour).Unix(), // 1 hour expire time
		})

	signed, _ := token.SignedString([]byte(key))

	return signed

}

func ParseToken(tokenString string) (*jwt.Token, error) {

	// decode
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("KEY")), nil
	})

	return token, err

}

func ValidateCurrentUser(id string, c *gin.Context) (models.User, bool) {
	// get current user
	// var oldUser models.User
	user, _ := c.Get("user")
	currentUser, _ := user.(models.User)
	success := true

	// validate authorization
	if cId := strconv.FormatUint(uint64(currentUser.ID), 10); id != cId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization"})
		success = false
	}

	return currentUser, success
}
