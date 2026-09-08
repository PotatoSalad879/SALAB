package routes

import (
	"github.com/gin-gonic/gin"

	"learnhub/internal/handlers"
)

// SetupRouter registers application routes and middleware.
func SetupRouter(
	healthHandler *handlers.HealthHandler,
	categoryHandler *handlers.CategoryHandler,
	courseHandler *handlers.CourseHandler,
) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", healthHandler.Index)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/categories", categoryHandler.GetAllCategories)
	router.GET("/courses", courseHandler.GetAllCourses)

	api := router.Group("/api/v1")
	{
		categories := api.Group("/categories")
		{
			categories.GET("", categoryHandler.GetAllCategories)
			categories.POST("", categoryHandler.CreateCategory)
		}

		courses := api.Group("/courses")
		{
			courses.GET("", courseHandler.GetAllCourses)
			courses.POST("", courseHandler.CreateCourse)
		}
	}

	return router
}
