package repository

import (
	"ecommerce/entity"
	"ecommerce/models"
	"encoding/json"
	"net/http"
)

type CategoryRestAPIDataSource struct {
}

// CreateCategory implements CategoryRepository.
func (c *CategoryRestAPIDataSource) CreateCategory(req models.CreateCategoryRequest) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryRepository.
func (c *CategoryRestAPIDataSource) DeleteCategory(id int) error {
	panic("unimplemented")
}

// GetCategoryList implements CategoryRepository.
func (c *CategoryRestAPIDataSource) GetCategoryList() ([]entity.Category, error) {
	result, err := http.Get("https://api.example.com/categories")

	if err != nil {
		return nil, err
	}

	jsonBody := json.NewDecoder(result.Body)

	defer result.Body.Close()

	var categories []entity.Category

	err = jsonBody.Decode(&categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory implements CategoryRepository.
func (c *CategoryRestAPIDataSource) UpdateCategory(req models.UpdateCategoryRequest, id int) error {
	panic("unimplemented")
}

func NewCategoryRestAPIRepository() CategoryRepository {
	return &CategoryRestAPIDataSource{}
}
