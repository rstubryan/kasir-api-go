package repository

import (
	"kasir-api/models"
	"sync"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(product models.Product) (*models.Product, error)
	Update(id int, product models.Product) (*models.Product, error)
	Delete(id int) error
}

type productRepository struct {
	products []models.Product
	mu       sync.RWMutex
	nextID   int
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		products: []models.Product{
			{ID: 1, Name: "Kopi Susu", Price: "15000", Stock: 50},
			{ID: 2, Name: "Nasi Goreng", Price: "25000", Stock: 30},
		},
		nextID: 3,
	}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.products, nil
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (r *productRepository) Create(product models.Product) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.nextID++

	r.products = append(r.products, product)
	return &product, nil
}

func (r *productRepository) Update(id int, product models.Product) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.products {
		if p.ID == id {
			product.ID = id
			r.products[i] = product
			return &r.products[i], nil
		}
	}
	return nil, nil
}

func (r *productRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return nil
}
