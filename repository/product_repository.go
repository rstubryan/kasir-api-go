package repository

import (
	"kasir-api/config"
	"kasir-api/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(id int, product *models.Product) error
	Delete(id int) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) Update(id int, product *models.Product) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id int) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
