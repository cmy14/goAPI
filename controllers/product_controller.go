package controllers

import (
	. "api-go/models"
	 "api-go/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: services.NewProductService(),
	}
}

func (ac *ProductController) GetAll(c *gin.Context) {
	// Appeler le service pour récupérer tous les Products
	products := ac.productService.GetAll()
	c.JSON(200, products)
}

func (ac *ProductController) Create(c *gin.Context) {
	var productModel Product
	if err := c.ShouldBindJSON(&productModel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Appeler le service pour créer un nouvel Product
	ac.productService.Create(productModel)
	c.JSON(201, gin.H{"message": "Product créé avec succès"})
}
