package service

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repository"

	"gorm.io/gorm"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(id int, category *models.Category) error
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return category, nil
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	return s.repo.Create(category)
}

func (s *categoryService) UpdateCategory(id int, category *models.Category) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	if category.Name == "" {
		return errors.New("category name is required")
	}

	return s.repo.Update(id, category)
}

func (s *categoryService) DeleteCategory(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
