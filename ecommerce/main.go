package main

import (
	"ecommerce/database"
	"ecommerce/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDatabase()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer sqlDB.Close()

	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	categoryHandler := handler.NewCategoryHandler(db)
	productHandler := handler.NewProductHandler(db)

	r.GET("/categories", categoryHandler.GetCategory)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.DetailCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	r.GET("/products", productHandler.GetProductList)
	r.POST("/products", productHandler.CreateProduct)

	r.Run(":8001")
}
