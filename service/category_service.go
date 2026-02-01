package service

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repository"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category models.Category) (*models.Category, error)
	UpdateCategory(id int, category models.Category) (*models.Category, error)
	DeleteCategory(id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (s *categoryService) CreateCategory(category models.Category) (*models.Category, error) {
	if category.Name == "" {
		return nil, errors.New("category name is required")
	}
	return s.repo.Create(category)
}

func (s *categoryService) UpdateCategory(id int, category models.Category) (*models.Category, error) {
	existingCategory, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, errors.New("category not found")
	}

	if category.Name == "" {
		return nil, errors.New("category name is required")
	}

	return s.repo.Update(id, category)
}

func (s *categoryService) DeleteCategory(id int) error {
	existingCategory, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return errors.New("category not found")
	}

	return s.repo.Delete(id)
}
