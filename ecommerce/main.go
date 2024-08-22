package main

import (
	"ecommerce/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	categoryHandler := handler.CategoryHandler{}
	productHandler := handler.ProductHandler{}

	r.GET("/categories", categoryHandler.GetCategory)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.DetailCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	r.GET("/products", productHandler.GetProductList)

	r.Run(":8001")
}
