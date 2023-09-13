package routes

import (
	"github.com/gin-gonic/gin"

	"api-go/controllers"
)

func InitializeRoutes(router *gin.Engine) {
	productController := controllers.NewProductController()
	apir := router.Group("api")
	{
		products := apir.Group("products")
		{
			products.GET("/", productController.GetAll)
			products.POST("/", productController.Create)
		}
	}

}
