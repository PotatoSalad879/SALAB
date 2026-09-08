package services

import (
	"errors"

	"learnhub/internal/dto"
	"learnhub/internal/models"
	"learnhub/internal/repositories"
)

// CourseService defines business operations on courses.
type CourseService interface {
	CreateCourse(input dto.CreateCourseRequest) (*dto.CourseResponse, error)
	GetAllCourses() ([]dto.CourseResponse, error)
	GetCourseByID(id uint) (*dto.CourseResponse, error)
}

type courseService struct {
	repo repositories.CourseRepository
}

// NewCourseService creates a new CourseService.
func NewCourseService(repo repositories.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(input dto.CreateCourseRequest) (*dto.CourseResponse, error) {
	course := &models.Course{
		Title:        input.Title,
		Slug:         input.Slug,
		Price:        input.Price,
		CategoryID:   input.CategoryID,
		InstructorID: input.InstructorID,
	}

	if err := s.repo.Create(course); err != nil {
		return nil, err
	}
	return mapCourseToResponse(course), nil
}

func (s *courseService) GetAllCourses() ([]dto.CourseResponse, error) {
	courses, err := s.repo.FetchAll()
	if err != nil {
		return nil, err
	}

	res := make([]dto.CourseResponse, 0, len(courses))
	for _, c := range courses {
		res = append(res, *mapCourseToResponse(&c))
	}
	return res, nil
}

func (s *courseService) GetCourseByID(id uint) (*dto.CourseResponse, error) {
	course, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, errors.New("course not found")
	}
	return mapCourseToResponse(course), nil
}

func mapCourseToResponse(c *models.Course) *dto.CourseResponse {
	return &dto.CourseResponse{
		ID:           c.ID,
		Title:        c.Title,
		Slug:         c.Slug,
		Price:        c.Price,
		CategoryID:   c.CategoryID,
		InstructorID: c.InstructorID,
		CreatedAt:    c.CreatedAt,
	}
}
