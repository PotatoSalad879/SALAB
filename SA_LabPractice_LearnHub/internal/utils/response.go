package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONSuccess sends a standardized successful JSON response.
func JSONSuccess(c *gin.Context, status int, data interface{}) {
	if status == http.StatusNoContent {
		c.Status(http.StatusNoContent)
		return
	}
	c.JSON(status, gin.H{
		"success": true,
		"data":    data,
	})
}

// JSONError sends a standardized error JSON response.
func JSONError(c *gin.Context, status int, message, detail string) {
	c.JSON(status, gin.H{
		"success": false,
		"error": gin.H{
			"message": message,
			"detail":  detail,
		},
	})
}

// SuccessResponse sends a standardized successful JSON response with a message.
func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	if status == http.StatusNoContent {
		c.Status(http.StatusNoContent)
		return
	}
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse sends a standardized error JSON response with a message.
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
}
