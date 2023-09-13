package models

import (
	"gorm.io/gorm"
	//"time"
)

type Product struct {
	gorm.Model
	Code            string `json:"code"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Quantity        int    `json:"quantite"`
	InventoryStatus string `json:"inventoryStatus"`
	Category        string `json:"category"`
	Image           string `json:"image"`
	Rating          int    `json:"rating"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

type ProductDTO struct {
	id              uint
	code            string
	name            string
	description     string
	price           int
	quantity        int
	inventoryStatus string
	category        string
	image           string
	rating          int
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

// func ProductTODTO(product Product) ProductDTO {
// 	productDto := ProductDTO{

// 		product.ID,
// 		product.Code,
// 		product.name,
// 		product.description,
// 		product.price,
// 		product.quantity,
// 		product.inventoryStatus,
// 		product.category,
// 		product.image,
// 		product.rating,
// 	}
// 	return productDto
// }
