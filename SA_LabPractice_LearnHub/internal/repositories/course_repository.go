package repositories

import (
	"errors"

	"learnhub/internal/models"

	"gorm.io/gorm"
)

// CourseRepository defines methods to access course data.
type CourseRepository interface {
	Create(course *models.Course) error
	FetchAll() ([]models.Course, error)
	FindByID(id uint) (*models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository creates a new CourseRepository instance.
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) FetchAll() ([]models.Course, error) {
	var courses []models.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) FindByID(id uint) (*models.Course, error) {
	var course models.Course
	if err := r.db.First(&course, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}
