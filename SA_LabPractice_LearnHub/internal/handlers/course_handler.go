package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learnhub/internal/dto"
	"learnhub/internal/services"
	"learnhub/internal/utils"
)

// CourseHandler handles course endpoints.
type CourseHandler struct {
	service services.CourseService
}

// NewCourseHandler constructs a new CourseHandler.
func NewCourseHandler(service services.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// GetAllCourses handles fetching all courses.
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.service.GetAllCourses()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "failed to load courses", err.Error())
		return
	}
	c.JSON(http.StatusOK, courses)
}

// CreateCourse handles creating a new course.
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req dto.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	course, err := h.service.CreateCourse(req)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "failed to create course", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, course)
}
