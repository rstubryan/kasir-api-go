package service

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repository"

	"gorm.io/gorm"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(id int, product *models.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id int) (*models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

func (s *productService) CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price == "" {
		return errors.New("product price is required")
	}
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id int, product *models.Product) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price == "" {
		return errors.New("product price is required")
	}

	return s.repo.Update(id, product)
}

func (s *productService) DeleteProduct(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
