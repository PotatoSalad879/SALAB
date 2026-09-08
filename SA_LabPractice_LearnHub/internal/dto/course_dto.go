package dto

import "time"

// CourseResponse defines the payload returned for course queries.
type CourseResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Price        float64   `json:"price"`
	CategoryID   uint      `json:"category_id"`
	InstructorID uint      `json:"instructor_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// CreateCourseRequest defines the payload for creating a new course.
type CreateCourseRequest struct {
	Title        string  `json:"title" binding:"required,min=2,max=200"`
	Slug         string  `json:"slug" binding:"required,min=2,max=200"`
	Price        float64 `json:"price" binding:"gte=0"`
	CategoryID   uint    `json:"category_id" binding:"required"`
	InstructorID uint    `json:"instructor_id" binding:"required"`
}
