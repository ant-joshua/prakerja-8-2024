package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" name:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" name:"name" binding:"required"`
}
