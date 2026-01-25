package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @description API Sederhana untuk aplikasi kasir dengan CRUD Products dan Categories
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https

type Product struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Kopi Susu"`
	Price string `json:"price" example:"15000"`
	Stock int    `json:"stock" example:"50"`
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
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Makanan"`
	Description string `json:"description" example:"Berbagai jenis makanan"`
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

// Health godoc
// @Summary Health check
// @Description Check if the server is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]

// GetAllProducts godoc
// @Summary Get all products
// @Description Get list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} Product
// @Router /api/v1/products [get]

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body Product true "Product object"
// @Success 201 {object} Product
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/products [post]

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a single product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 404 {string} string "Product not found"
// @Router /api/v1/products/{id} [get]

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Product object"
// @Success 200 {object} Product
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Product not found"
// @Router /api/v1/products/{id} [put]

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {string} string "Product not found"
// @Router /api/v1/products/{id} [delete]

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get list of all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} Category
// @Router /api/v1/categories [get]

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body Category true "Category object"
// @Success 201 {object} Category
// @Failure 400 {string} string "Bad Request"
// @Router /api/v1/categories [post]

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Get a single category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} Category
// @Failure 404 {string} string "Category not found"
// @Router /api/v1/categories/{id} [get]

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body Category true "Category object"
// @Success 200 {object} Category
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Category not found"
// @Router /api/v1/categories/{id} [put]

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204
// @Failure 404 {string} string "Category not found"
// @Router /api/v1/categories/{id} [delete]

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

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server is running on port 8080")
	fmt.Println("Swagger UI available at: http://localhost:8080/swagger/index.html")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
