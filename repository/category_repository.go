package repository

import (
	"kasir-api/config"
	"kasir-api/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(category *models.Category) error
	Update(id int, category *models.Category) error
	Delete(id int) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	var category models.Category
	err := config.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	return config.DB.Create(category).Error
}

func (r *categoryRepository) Update(id int, category *models.Category) error {
	return config.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
}

func (r *categoryRepository) Delete(id int) error {
	return config.DB.Delete(&models.Category{}, id).Error
}
