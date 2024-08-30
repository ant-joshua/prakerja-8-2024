package service

import (
	"ecommerce/models"

	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(req models.OrderCreateRequest, userID int) (*models.Order, error)
}

type OrderServiceImpl struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &OrderServiceImpl{
		db: db,
	}
}

func (o *OrderServiceImpl) CreateOrder(req models.OrderCreateRequest, userID int) (*models.Order, error) {

	// 1. Create order
	order := models.Order{
		UserID:    uint(userID),
		OrderCode: "ORDER-123",
	}

	// 2. Create order items
	var orderItems []models.OrderItems
	var products []models.Product

	var productIDS []int

	for _, item := range req.Items {
		productIDS = append(productIDS, item.ProductID)

		orderItem := models.OrderItems{
			ProductID: item.ProductID,
			Qty:       item.Qty,
			Price:     0,
		}

		orderItems = append(orderItems, orderItem)
	}

	err := o.db.Model(&models.Product{}).Where("id IN ?", productIDS).Find(&products).Error

	orderSubTotal := 0
	orderTotal := 0

	for i, item := range orderItems {
		orderItems[i].Price = products[i].Price
		orderItems[i].SubTotal = products[i].Price * item.Qty

		orderSubTotal += orderItems[i].SubTotal
	}

	order.Discount = req.Discount
	order.Tax = req.Tax
	order.SubTotal = orderSubTotal

	orderTotal = (orderSubTotal - req.Discount) + req.Tax

	order.Total = orderTotal

	if err != nil {
		return nil, err
	}

	order.Items = orderItems

	// 3. Save to database
	err = o.db.Create(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil

}
