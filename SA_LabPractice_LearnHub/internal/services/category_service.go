package services

import (
	"errors"

	"learnhub/internal/dto"
	"learnhub/internal/models"
	"learnhub/internal/repositories"
)

// CategoryService defines business operations on categories.
type CategoryService interface {
	CreateCategory(input dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetCategoryByID(id uint) (*dto.CategoryResponse, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

// NewCategoryService creates a new CategoryService.
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(input dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := &models.Category{
		Name: input.Name,
		Slug: input.Slug,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return mapCategoryToResponse(category), nil
}

func (s *categoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.FetchAll()
	if err != nil {
		return nil, err
	}

	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, c := range categories {
		res = append(res, *mapCategoryToResponse(&c))
	}
	return res, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return mapCategoryToResponse(category), nil
}

func mapCategoryToResponse(c *models.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:        c.ID,
		Name:      c.Name,
		Slug:      c.Slug,
		CreatedAt: c.CreatedAt,
	}
}
