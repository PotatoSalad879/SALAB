package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles system status and root index routes.
type HealthHandler struct{}

// NewHealthHandler constructs a HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Index handles root route request.
func (h *HealthHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "LearnHub - Online Course Platform API (GORM + PostgreSQL)",
		"endpoints": []string{
			"GET /health",
			"GET /categories",
			"GET /courses",
			"GET /api/v1/categories",
			"GET /api/v1/courses",
		},
	})
}

// HealthCheck handles health check request.
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "connected",
	})
}
