package main

import (
	"github.com/gin-gonic/gin"
    "api-go/routes"
   . "api-go/utils"
    
)

func main() {
	router := gin.Default()
    InitializeDatabase()
    routes.InitializeRoutes(router)
   
	router.Run("localhost:8080")
}
