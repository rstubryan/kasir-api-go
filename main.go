package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/config"
	"kasir-api/handlers"
	"kasir-api/migrations"
	"kasir-api/repository"
	"kasir-api/service"
	"log"
	"net/http"

	_ "kasir-api/docs"

	"github.com/gorilla/mux"
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

// Health godoc
// @Summary Health check
// @Description Check if the server is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Service is healthy",
	})
}

func main() {
	// Initialize database
	err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	// Run migrations
	migrations.RunMigrations()

	// Seed initial data
	migrations.SeedData()

	// Initialize repositories
	productRepo := repository.NewProductRepository()
	categoryRepo := repository.NewCategoryRepository()

	// Initialize services
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Create router
	router := mux.NewRouter()

	// Register routes
	productHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)

	// Health check endpoint
	router.HandleFunc("/health", healthHandler).Methods("GET")

	// Swagger documentation with dynamic host
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	fmt.Println("Server is running on port 8080")
	fmt.Println("Swagger UI available at: http://localhost:8080/swagger/index.html")
	fmt.Println("Connected to NeonDB database")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
