package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidationError(err error, c *gin.Context) {

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}

func RecordNotFoundError(result *gorm.DB, c *gin.Context) {

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Record not found
		// Handle this case
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
	} else {
		// Other error occurred
		// Handle the error
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

}
