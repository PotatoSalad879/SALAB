package repositories

import (
	"context"

	"learnhub/internal/models"

	"gorm.io/gorm"
)

// Store provides direct database operations across all models.
type Store struct {
	db *gorm.DB
}

// New creates a new Store instance.
func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

// DB returns the underlying gorm.DB.
func (s *Store) DB() *gorm.DB {
	return s.db
}

func (s *Store) CreateUser(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *Store) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.db.WithContext(ctx).Create(category).Error
}

func (s *Store) CreateCourse(ctx context.Context, course *models.Course) error {
	return s.db.WithContext(ctx).Create(course).Error
}

func (s *Store) CreateTag(ctx context.Context, tag *models.Tag) error {
	return s.db.WithContext(ctx).Create(tag).Error
}

func (s *Store) CreateEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return s.db.WithContext(ctx).Create(enrollment).Error
}

func (s *Store) CreateLesson(ctx context.Context, lesson *models.Lesson) error {
	return s.db.WithContext(ctx).Create(lesson).Error
}

func (s *Store) CreateReview(ctx context.Context, review *models.Review) error {
	return s.db.WithContext(ctx).Create(review).Error
}

func (s *Store) CreateCourseSetting(ctx context.Context, setting *models.CourseSetting) error {
	return s.db.WithContext(ctx).Create(setting).Error
}

func (s *Store) FetchAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	if err := s.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *Store) FetchAllCourses(ctx context.Context) ([]models.Course, error) {
	var courses []models.Course
	if err := s.db.WithContext(ctx).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
