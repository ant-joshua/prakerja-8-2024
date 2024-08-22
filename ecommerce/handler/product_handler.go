package handler

import (
	"ecommerce/helpers"
	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func (h *ProductHandler) GetProductList(c *gin.Context) {

	productList := []models.Product{
		{
			ID:         1,
			Name:       "Laptop",
			Category:   models.Category{ID: 1, Name: "Electronic"},
			CategoryID: 1,
			Price:      10000000,
			Stock:      10,
		},
	}

	c.JSON(200, helpers.NewSuccessResponse(productList))
}
