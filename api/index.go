package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	products   = []Product{}
	categories = []Category{}
	once       sync.Once
)

func initData() {
	products = []Product{
		{
			ID:    1,
			Name:  "Kopi Susu",
			Price: "15000",
			Stock: 50,
		},
		{
			ID:    2,
			Name:  "Nasi Goreng",
			Price: "25000",
			Stock: 30,
		},
	}

	categories = []Category{
		{
			ID:          1,
			Name:        "Makanan",
			Description: "Berbagai jenis makanan",
		},
		{
			ID:          2,
			Name:        "Minuman",
			Description: "Berbagai jenis minuman",
		},
	}
}

func initApp() {
	initData()

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Service is healthy",
		})
	})

	http.HandleFunc("GET /api/v1/products", getAllProducts)
	http.HandleFunc("POST /api/v1/products", createProduct)
	http.HandleFunc("GET /api/v1/products/", getProductByID)
	http.HandleFunc("PUT /api/v1/products/", updateProduct)
	http.HandleFunc("DELETE /api/v1/products/", deleteProduct)

	http.HandleFunc("GET /api/v1/categories", getAllCategories)
	http.HandleFunc("POST /api/v1/categories", createCategory)
	http.HandleFunc("GET /api/v1/categories/", getCategoryByID)
	http.HandleFunc("PUT /api/v1/categories/", updateCategory)
	http.HandleFunc("DELETE /api/v1/categories/", deleteCategory)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, product := range products {
		if product.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedProduct.ID = id
	products[index] = updatedProduct
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products[index])
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, product := range products {
		if product.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	products = append(products[:index], products[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCategory)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for _, category := range categories {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, category := range categories {
		if category.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	var updatedCategory Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedCategory.ID = id
	categories[index] = updatedCategory
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories[index])
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, category := range categories {
		if category.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	categories = append(categories[:index], categories[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initApp)
	http.DefaultServeMux.ServeHTTP(w, r)
}
