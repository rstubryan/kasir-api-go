package migrations

import (
	"kasir-api/config"
	"kasir-api/models"
	"log"
)

func RunMigrations() {
	db := config.GetDB()

	// Auto migrate tables
	err := db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Transaction{}, &models.TransactionDetail{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
}

func SeedData() {
	db := config.GetDB()

	// Check if categories already exist
	var categoryCount int64
	db.Model(&models.Category{}).Count(&categoryCount)
	if categoryCount == 0 {
		categories := []models.Category{
			{Name: "Makanan", Description: "Berbagai jenis makanan"},
			{Name: "Minuman", Description: "Berbagai jenis minuman"},
		}
		for _, category := range categories {
			db.Create(&category)
		}
		log.Println("Seeded categories data")
	}

	// Check if products already exist
	var productCount int64
	db.Model(&models.Product{}).Count(&productCount)
	if productCount == 0 {
		products := []models.Product{
			{Name: "Kopi Susu", Price: "15000", Stock: 50},
			{Name: "Nasi Goreng", Price: "25000", Stock: 30},
		}
		for _, product := range products {
			db.Create(&product)
		}
		log.Println("Seeded products data")
	}
}
