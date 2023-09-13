package services

import (
	. "api-go/models"

	"api-go/utils"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}
func (as *ProductService) GetAll() []Product {
	var products []Product
	utils.DB.Find(&products)
	return products
}

func (as *ProductService) Create(product Product) error {
	result := utils.DB.Create(&product)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
