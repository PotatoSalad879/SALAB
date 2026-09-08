package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"learnhub/internal/config"
	"learnhub/internal/handlers"
	"learnhub/internal/repositories"
	"learnhub/internal/routes"
	"learnhub/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading configuration from environment")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database handle failed: %v", err)
	}
	defer sqlDB.Close()

	log.Println("connected to PostgreSQL and database migrated successfully via GORM")

	categoryRepo := repositories.NewCategoryRepository(db)
	courseRepo := repositories.NewCourseRepository(db)

	categoryService := services.NewCategoryService(categoryRepo)
	courseService := services.NewCourseService(courseRepo)

	healthHandler := handlers.NewHealthHandler()
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	courseHandler := handlers.NewCourseHandler(courseService)

	router := routes.SetupRouter(healthHandler, categoryHandler, courseHandler)

	port := cfg.ServerPort
	if port == "" {
		port = "8000"
	}

	log.Printf("Server is starting on port %s (http://localhost:%s)\n", port, port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
