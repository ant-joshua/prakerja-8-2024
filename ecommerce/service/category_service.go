package service

import (
	"ecommerce/entity"
	"ecommerce/models"
	"ecommerce/repository"
)

type CategoryService interface {
	GetCategoryList() ([]entity.Category, error)
	CreateCategory(req models.CreateCategoryRequest) error
	UpdateCategory(req models.UpdateCategoryRequest, id int) error
	DeleteCategory(id int) error
}

type CategoryServiceImpl struct {
	categoryRepo repository.CategoryRepository
}

// CreateCategory implements CategoryService.
func (c *CategoryServiceImpl) CreateCategory(req models.CreateCategoryRequest) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryService.
func (c *CategoryServiceImpl) DeleteCategory(id int) error {
	panic("unimplemented")
}

// GetCategoryList implements CategoryService.
func (c *CategoryServiceImpl) GetCategoryList() ([]entity.Category, error) {
	categories, err := c.categoryRepo.GetCategoryList()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory implements CategoryService.
func (c *CategoryServiceImpl) UpdateCategory(req models.UpdateCategoryRequest, id int) error {
	panic("unimplemented")
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepo: categoryRepo,
	}
}
