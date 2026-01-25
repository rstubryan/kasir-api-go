package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Stock int    `json:"stock"`
}

var products = []Product{
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

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
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

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllProducts(w, r)
	case "POST":
		createProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProductByID(w, r)
	case "PUT":
		updateProduct(w, r)
	case "DELETE":
		deleteProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllCategories(w, r)
	case "POST":
		createCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func categoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCategoryByID(w, r)
	case "PUT":
		updateCategory(w, r)
	case "DELETE":
		deleteCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/v1/products", productsHandler)
	http.HandleFunc("/api/v1/products/", productByIDHandler)

	http.HandleFunc("/api/v1/categories", categoriesHandler)
	http.HandleFunc("/api/v1/categories/", categoryByIDHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK", "message": "Service is healthy"})
	})

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
