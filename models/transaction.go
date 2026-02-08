package models

import "time"

type Transaction struct {
	ID          int                 `json:"id" gorm:"primaryKey"`
	TotalAmount int                 `json:"total_amount" gorm:"not null"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details" gorm:"foreignKey:TransactionID"`
}

type TransactionDetail struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	TransactionID int    `json:"transaction_id" gorm:"not null"`
	ProductID     int    `json:"product_id" gorm:"not null"`
	ProductName   string `json:"product_name,omitempty" gorm:"-"`
	Quantity      int    `json:"quantity" gorm:"not null"`
	Subtotal      int    `json:"subtotal" gorm:"not null"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type DailyReport struct {
	TotalRevenue    int                    `json:"total_revenue"`
	TotalTransaction int                   `json:"total_transaksi"`
	BestSeller      BestSellerProduct      `json:"produk_terlaris"`
}

type BestSellerProduct struct {
	Name     string `json:"nama"`
	QtySold  int    `json:"qty_terjual"`
}
