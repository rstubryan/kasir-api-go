package repository

import (
	"kasir-api/config"
	"kasir-api/models"

	"errors"
)

type TransactionRepository interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
	GetDailyReport() (*models.DailyReport, error)
	GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

// CreateTransaction creates a new transaction with its details
// FIX: Properly insert transaction details with correct transaction ID
func (r *transactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx := config.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	// Process each item
	for _, item := range items {
		var product models.Product

		// Get product details
		if err := tx.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("product not found")
		}

		// Check stock
		if product.Stock < item.Quantity {
			tx.Rollback()
			return nil, errors.New("insufficient stock")
		}

		// Calculate subtotal
		subtotal := product.PriceInt() * item.Quantity
		totalAmount += subtotal

		// Update product stock
		if err := tx.Model(&product).Update("stock", product.Stock-item.Quantity).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Create transaction detail
		detail := models.TransactionDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Subtotal:  subtotal,
		}
		details = append(details, detail)
	}

	// Create transaction
	transaction := models.Transaction{
		TotalAmount: totalAmount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// FIX: Insert transaction details with correct transaction ID
	// The issue in the original code was that it was modifying details slice in place
	// without properly associating it with the created transaction
	for i := range details {
		details[i].TransactionID = transaction.ID
		details[i].ProductName = "" // Ensure this is not saved to DB

		if err := tx.Create(&details[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Load details with product names for response
	var result models.Transaction
	if err := config.DB.Preload("Details").First(&result, transaction.ID).Error; err != nil {
		return nil, err
	}

	// Manually populate product names
	for i := range result.Details {
		var product models.Product
		config.DB.First(&product, result.Details[i].ProductID)
		result.Details[i].ProductName = product.Name
	}

	return &result, nil
}

// GetDailyReport gets sales summary for today
func (r *transactionRepository) GetDailyReport() (*models.DailyReport, error) {
	var report models.DailyReport

	// Get total revenue today
	if err := config.DB.Table("transactions").
		Where("DATE(created_at) = CURRENT_DATE").
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&report.TotalRevenue).Error; err != nil {
		return nil, err
	}

	// Get total transaction count today
	if err := config.DB.Table("transactions").
		Where("DATE(created_at) = CURRENT_DATE").
		Select("COUNT(*)").
		Scan(&report.TotalTransaction).Error; err != nil {
		return nil, err
	}

	// Get best seller product
	type BestSellerResult struct {
		ProductName string
		TotalQty    int
	}
	var bestSeller BestSellerResult

	if err := config.DB.Table("transaction_details").
		Select("p.name as product_name, SUM(transaction_details.quantity) as total_qty").
		Joins("JOIN products p ON p.id = transaction_details.product_id").
		Joins("JOIN transactions t ON t.id = transaction_details.transaction_id").
		Where("DATE(t.created_at) = CURRENT_DATE").
		Group("p.name").
		Order("total_qty DESC").
		Limit(1).
		Scan(&bestSeller).Error; err != nil {
		return nil, err
	}

	if bestSeller.ProductName != "" {
		report.BestSeller = models.BestSellerProduct{
			Name:    bestSeller.ProductName,
			QtySold: bestSeller.TotalQty,
		}
	}

	return &report, nil
}

// GetReportByDateRange gets sales summary for a date range
func (r *transactionRepository) GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error) {
	var report models.DailyReport

	// Get total revenue for date range
	if err := config.DB.Table("transactions").
		Where("DATE(created_at) >= ? AND DATE(created_at) <= ?", startDate, endDate).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&report.TotalRevenue).Error; err != nil {
		return nil, err
	}

	// Get total transaction count for date range
	if err := config.DB.Table("transactions").
		Where("DATE(created_at) >= ? AND DATE(created_at) <= ?", startDate, endDate).
		Select("COUNT(*)").
		Scan(&report.TotalTransaction).Error; err != nil {
		return nil, err
	}

	// Get best seller product
	type BestSellerResult struct {
		ProductName string
		TotalQty    int
	}
	var bestSeller BestSellerResult

	if err := config.DB.Table("transaction_details").
		Select("p.name as product_name, SUM(transaction_details.quantity) as total_qty").
		Joins("JOIN products p ON p.id = transaction_details.product_id").
		Joins("JOIN transactions t ON t.id = transaction_details.transaction_id").
		Where("DATE(t.created_at) >= ? AND DATE(t.created_at) <= ?", startDate, endDate).
		Group("p.name").
		Order("total_qty DESC").
		Limit(1).
		Scan(&bestSeller).Error; err != nil {
		return nil, err
	}

	if bestSeller.ProductName != "" {
		report.BestSeller = models.BestSellerProduct{
			Name:    bestSeller.ProductName,
			QtySold: bestSeller.TotalQty,
		}
	}

	return &report, nil
}
