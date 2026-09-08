package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learnhub/internal/dto"
	"learnhub/internal/services"
	"learnhub/internal/utils"
)

// CategoryHandler handles category endpoints.
type CategoryHandler struct {
	service services.CategoryService
}

// NewCategoryHandler constructs a new CategoryHandler.
func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetAllCategories handles fetching all categories.
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "failed to load categories", err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory handles creating a new category.
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	category, err := h.service.CreateCategory(req)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "failed to create category", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, category)
}
