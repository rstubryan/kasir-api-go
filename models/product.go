package models

type Product struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Kopi Susu"`
	Price string `json:"price" example:"15000"`
	Stock int    `json:"stock" example:"50"`
}
