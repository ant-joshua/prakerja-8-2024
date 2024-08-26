package main

import (
	"ecommerce/database"
	"ecommerce/handler"
	"ecommerce/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	authHandler := handler.NewAuthHandler(db)

	r.GET("/categories", middleware.AuthMiddleware(), categoryHandler.GetCategory)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.DetailCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	r.GET("/products", productHandler.GetProductList)
	r.POST("/products", productHandler.CreateProduct)

	// Auth handler
	r.POST("/register", authHandler.Register)
	r.POST("/verify-otp", authHandler.VerifyOTP)
	r.POST("/login", authHandler.Login)

	r.Run(":8001")
}
