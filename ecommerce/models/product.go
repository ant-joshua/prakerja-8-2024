package models

type Product struct {
	ID         int      `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	CategoryID int      `json:"category_id"`
	Price      int      `json:"price"`
	Stock      int      `json:"stock"`
}

type CreateProductRequest struct {
	Name       string `json:"name" name:"name" binding:"required"`
	CategoryID int    `json:"category_id" name:"category_id" binding:"required"`
	Price      int    `json:"price" name:"price" binding:"required"`
	Stock      int    `json:"stock" name:"stock" binding:"required"`
}
