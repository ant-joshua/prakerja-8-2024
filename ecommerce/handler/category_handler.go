package handler

import (
	"ecommerce/helpers"
	"ecommerce/models"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var categoryList = []models.Category{
	{
		ID:   1,
		Name: "Electronic",
	},
	{
		ID:   2,
		Name: "Fashion",
	},
	{
		ID:   3,
		Name: "Food",
	},
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": categoryList,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.CreateCategoryRequest

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	c.JSON(200, gin.H{
		"data": category,
	})
}

func (h *CategoryHandler) DetailCategory(c *gin.Context) {
	categoryID := c.Param("id")

	parseCategoryID, err := strconv.Atoi(categoryID)

	if err != nil {
		c.JSON(400, helpers.NewErrorResponse[any](400, "Invalid category ID"))
		return
	}

	var categoryData *models.Category

	for _, category := range categoryList {
		if category.ID == parseCategoryID {
			categoryData = &category
			break
		}
	}

	if categoryData == nil {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	categoryData.ID = parseCategoryID

	c.JSON(200, gin.H{
		"data": &categoryData,
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	parseCategoryID, err := strconv.Atoi(categoryID)

	if err != nil {
		c.JSON(400, helpers.NewErrorResponse[any](400, "Invalid category ID"))
		return
	}

	var categoryData *models.Category
	var categoryIndex int

	for index, category := range categoryList {
		if category.ID == parseCategoryID {
			categoryIndex = index
			categoryData = &category
			break
		}
	}

	if categoryData == nil {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	var updateCategory models.UpdateCategoryRequest

	err = c.BindJSON(&updateCategory)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	categoryData.Name = updateCategory.Name

	categoryList[categoryIndex] = *categoryData

	c.JSON(200, gin.H{
		"data": &categoryData,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	parseCategoryID, err := strconv.Atoi(categoryID)

	if err != nil {
		c.JSON(400, helpers.NewErrorResponse[any](400, "Invalid category ID"))
		return
	}

	var categoryData *models.Category
	var categoryIndex int

	for index, category := range categoryList {
		if category.ID == parseCategoryID {
			categoryIndex = index
			categoryData = &category
			break
		}
	}

	if categoryData == nil {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	categoryList = slices.Delete(categoryList, categoryIndex, 1)

	c.JSON(200, gin.H{
		"message": "Category deleted",
	})
}
