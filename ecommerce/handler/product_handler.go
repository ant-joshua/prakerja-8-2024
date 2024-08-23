package handler

import (
	"ecommerce/helpers"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

// CreateProduct implements ProductHandlerInterface.
func (h *ProductHandler) CreateProduct(c *gin.Context) {

	var req models.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Failed to create product", err))
		return
	}

	// Create product
	product := models.Product{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	err := h.db.Create(&product).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to create product"))
		return
	}

	c.JSON(200, helpers.NewSuccessResponse("Product created"))
}

type ProductHandlerInterface interface {
	GetProductList(c *gin.Context)
	CreateProduct(c *gin.Context)
}

func NewProductHandler(db *gorm.DB) ProductHandlerInterface {
	return &ProductHandler{
		db: db,
	}
}

func (h *ProductHandler) GetProductList(c *gin.Context) {

	var productList []models.Product

	err := h.db.Preload("Category").Find(&productList).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to fetch product"))
		return
	}

	c.JSON(200, helpers.NewSuccessResponse(productList))
}
