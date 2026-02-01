package models

type Category struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Makanan"`
	Description string `json:"description" example:"Berbagai jenis makanan"`
}
