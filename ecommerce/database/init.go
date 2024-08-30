package database

import (
	"ecommerce/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dbPath string) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{}, &models.Order{},
		&models.OrderItems{}, &models.Permission{}, &models.Role{}, &models.RolePermission{})

	return db, nil
}
