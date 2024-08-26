package handler

import (
	"ecommerce/helpers"
	"ecommerce/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

// AssignProductToCategory implements CategoryHandlerInterface.
func (h *CategoryHandler) AssignProductToCategory(c *gin.Context) {
	panic("unimplemented")
}

type CategoryHandlerInterface interface {
	GetCategory(c *gin.Context)
	CreateCategory(c *gin.Context)
	DetailCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	AssignProductToCategory(c *gin.Context)
}

func NewCategoryHandler(db *gorm.DB) CategoryHandlerInterface {
	return &CategoryHandler{
		db: db,
	}
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

	user := c.MustGet("user").(jwt.MapClaims)

	fmt.Printf("User: %+v\n", user)
	userId := user["id"].(float64)

	fmt.Printf("User ID: %v\n", userId)

	var categories []models.Category

	err := h.db.Find(&categories).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to fetch category"))
		return
	}

	c.JSON(200, gin.H{
		"data": categories,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	category := models.Category{
		Name: req.Name,
	}

	err = h.db.Create(&category).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to create category"))
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

	err = h.db.Where("id = ?", parseCategoryID).First(&categoryData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Internal Server Error"))
		return
	}

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

	err = h.db.Where("id = ?", parseCategoryID).First(&categoryData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Internal Server Error"))
		return
	}

	var updateCategory models.UpdateCategoryRequest

	err = c.BindJSON(&updateCategory)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	categoryData.Name = updateCategory.Name

	err = h.db.Save(&categoryData).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to update category"))
		return
	}

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

	err = h.db.Where("id = ?", parseCategoryID).First(&categoryData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, helpers.NewErrorResponse[any](404, "Category not found"))
		return
	}

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Internal Server Error"))
		return
	}

	err = h.db.Delete(&categoryData).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to delete category"))
		return
	}

	c.JSON(200, gin.H{
		"message": "Category deleted",
	})
}
