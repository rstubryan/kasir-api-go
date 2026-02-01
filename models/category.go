package models

import "time"

type Category struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Makanan"`
	Description string    `json:"description" example:"Berbagai jenis makanan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
