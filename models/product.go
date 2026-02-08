package models

import (
	"strconv"
	"time"
)

type Product struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Kopi Susu"`
	Price     string    `json:"price" example:"15000"`
	Stock     int       `json:"stock" example:"50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PriceInt converts Price string to int
func (p *Product) PriceInt() int {
	price, _ := strconv.Atoi(p.Price)
	return price
}
