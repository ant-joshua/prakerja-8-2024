package service_test

import (
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/service"
	"log"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	database, err := database.InitDatabase("../test.db")

	db = database

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer sqlDB.Close()

	m.Run()
}

func TestCreateOrder(t *testing.T) {
	req := models.OrderCreateRequest{
		Discount: 10000,
		Tax:      1000,
		Items: []models.OrderItemCreateRequest{
			{
				ProductID: 1, // harga nya 1000 an per item
				Qty:       20,
			},
		},
	}

	orderService := service.NewOrderService(db)

	result, err := orderService.CreateOrder(req, 1)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result.OrderCode == "" {
		t.Errorf("Error: Order code is empty")
	}

	if result.Total == 0 {
		t.Errorf("Error: Total is 0")
	}

	if len(result.Items) == 0 {
		t.Errorf("Error: Items is empty")
	}

	if result.Total != 11000 {
		t.Errorf("Error: Total is not correct")
	}
}
