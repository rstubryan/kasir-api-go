package models

import "time"

type Product struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Kopi Susu"`
	Price     string    `json:"price" example:"15000"`
	Stock     int       `json:"stock" example:"50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
