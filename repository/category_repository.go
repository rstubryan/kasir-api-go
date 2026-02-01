package repository

import (
	"kasir-api/models"
	"sync"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(category models.Category) (*models.Category, error)
	Update(id int, category models.Category) (*models.Category, error)
	Delete(id int) error
}

type categoryRepository struct {
	categories []models.Category
	mu         sync.RWMutex
	nextID     int
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		categories: []models.Category{
			{ID: 1, Name: "Makanan", Description: "Berbagai jenis makanan"},
			{ID: 2, Name: "Minuman", Description: "Berbagai jenis minuman"},
		},
		nextID: 3,
	}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.categories, nil
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, category := range r.categories {
		if category.ID == id {
			return &category, nil
		}
	}
	return nil, nil
}

func (r *categoryRepository) Create(category models.Category) (*models.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	category.ID = r.nextID
	r.nextID++

	r.categories = append(r.categories, category)
	return &category, nil
}

func (r *categoryRepository) Update(id int, category models.Category) (*models.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, c := range r.categories {
		if c.ID == id {
			category.ID = id
			r.categories[i] = category
			return &r.categories[i], nil
		}
	}
	return nil, nil
}

func (r *categoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, category := range r.categories {
		if category.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return nil
}
