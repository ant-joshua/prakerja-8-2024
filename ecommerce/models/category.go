package models

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" name:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" name:"name" binding:"required"`
}

type CategoryPostgres struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CategoryMongo struct {
	ID   int    `json:"id" bson:"_id"`
	Nama string `json:"nama" bson:"nama"`
}
