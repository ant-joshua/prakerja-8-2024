package main

import (
	"ecommerce/database"
	"ecommerce/handler"
	"ecommerce/middleware"
	"ecommerce/repository"
	"ecommerce/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	db, err := database.InitDatabase("ecommerce.db")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer sqlDB.Close()

	redisCache := database.NewRedis(0)

	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// categoryRepo := repository.NewCategoryORMRepository(db) // gorm
	categoryRepo := repository.NewCategoryRestAPIRepository()
	categoryService := service.NewCategoryService(categoryRepo)

	categoryHandler := handler.NewCategoryHandler(db, categoryService)

	productHandler := handler.NewProductHandler(db)
	authHandler := handler.NewAuthHandler(db)

	authMiddleware := middleware.AuthMiddleware(db)
	permissionMiddleware := middleware.NewPermissionMiddleware(db, redisCache)

	r.GET("/categories", authMiddleware, categoryHandler.GetCategory)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.DetailCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	r.GET("/products", authMiddleware, permissionMiddleware.PermissionMiddlewareMap("read_product"), productHandler.GetProductList)
	r.POST("/products", authMiddleware, permissionMiddleware.PermissionMiddleware("create_product"), productHandler.CreateProduct)

	// Auth handler
	r.POST("/register", authHandler.Register)
	r.POST("/verify-otp", authHandler.VerifyOTP)
	r.POST("/login", authHandler.Login)

	r.Run(":8001")
}
