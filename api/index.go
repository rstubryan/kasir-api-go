package handler

import (
	"encoding/json"
	"kasir-api/config"
	"kasir-api/handlers"
	"kasir-api/migrations"
	"kasir-api/repository"
	"kasir-api/service"
	"net/http"
	"os"

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

// @BasePath /
// @schemes http https

var router *mux.Router

func init() {
	// Initialize database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback to other env vars for Vercel
		dsn = os.Getenv("POSTGRES_URL")
	}

	if dsn == "" {
		panic("DATABASE_URL or POSTGRES_URL environment variable is required")
	}

	err := config.InitDB()
	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

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
	router = mux.NewRouter()

	// Register routes
	productHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)

	// Health check endpoint
	router.HandleFunc("/health", healthHandler).Methods("GET")

	// Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Service is healthy",
	})
}

// Handler godoc
// @Summary Main handler for Vercel
// @Description This is the main entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	router.ServeHTTP(w, r)
}
