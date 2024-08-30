package database

import (
	"ecommerce/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{}, &models.Order{},
		&models.OrderItems{}, &models.Permission{}, &models.Role{}, &models.RolePermission{})

	return db, nil
}
