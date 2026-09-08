package repositories

import (
	"errors"

	"learnhub/internal/models"

	"gorm.io/gorm"
)

// CategoryRepository defines methods to access category data.
type CategoryRepository interface {
	Create(category *models.Category) error
	FetchAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new CategoryRepository instance.
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FetchAll() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}
