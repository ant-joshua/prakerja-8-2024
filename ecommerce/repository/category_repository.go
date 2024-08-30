package repository

import (
	"ecommerce/entity"
	"ecommerce/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoryList() ([]entity.Category, error)
	CreateCategory(req models.CreateCategoryRequest) error
	UpdateCategory(req models.UpdateCategoryRequest, id int) error
	DeleteCategory(id int) error
}

type CategoryDataORMSource struct {
	db *gorm.DB
}

// CreateCategory implements CategoryRepositoryInterface.
func (c *CategoryDataORMSource) CreateCategory(req models.CreateCategoryRequest) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryRepositoryInterface.
func (c *CategoryDataORMSource) DeleteCategory(id int) error {
	panic("unimplemented")
}

// GetCategoryList implements CategoryRepositoryInterface.
func (c *CategoryDataORMSource) GetCategoryList() ([]entity.Category, error) {
	var categoriesModel []models.Category

	err := c.db.Find(&categoriesModel).Error

	if err != nil {
		return nil, err
	}

	var categories []entity.Category

	for _, category := range categoriesModel {
		categories = append(categories, entity.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categories, nil
}

// UpdateCategory implements CategoryRepositoryInterface.
func (c *CategoryDataORMSource) UpdateCategory(req models.UpdateCategoryRequest, id int) error {
	panic("unimplemented")
}

func NewCategoryORMRepository(db *gorm.DB) CategoryRepository {
	return &CategoryDataORMSource{
		db: db,
	}
}
