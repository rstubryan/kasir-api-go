package service

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product models.Product) (*models.Product, error)
	UpdateProduct(id int, product models.Product) (*models.Product, error)
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
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productService) CreateProduct(product models.Product) (*models.Product, error) {
	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.Price == "" {
		return nil, errors.New("product price is required")
	}
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id int, product models.Product) (*models.Product, error) {
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.Price == "" {
		return nil, errors.New("product price is required")
	}

	return s.repo.Update(id, product)
}

func (s *productService) DeleteProduct(id int) error {
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existingProduct == nil {
		return errors.New("product not found")
	}

	return s.repo.Delete(id)
}
