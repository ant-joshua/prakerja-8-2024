package models

type Order struct {
	ID        int          `gorm:"primaryKey" json:"id"`
	OrderCode string       `gorm:"order_code" json:"order_code"`
	UserID    uint         `json:"user_id"`
	User      User         `gorm:"foreignKey:UserID" json:"user"`
	Items     []OrderItems `gorm:"foreignKey:OrderID" json:"items"`
	Discount  int          `gorm:"discount" json:"discount"`
	Tax       int          `gorm:"tax" json:"tax"`
	SubTotal  int          `gorm:"sub_total" json:"sub_total"`
	Total     int          `gorm:"total" json:"total"`
}

type OrderItems struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"foreignKey:OrderID" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	ProductID int     `gorm:"foreignKey:ProductID" json:"product_id"`
	Qty       int     `gorm:"qty" json:"qty"`
	Price     int     `gorm:"price" json:"price"`
	SubTotal  int     `gorm:"sub_total" json:"sub_total"`
}

type OrderItemCreateRequest struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

type OrderCreateRequest struct {
	Discount int                      `json:"discount"`
	Tax      int                      `json:"tax"`
	Items    []OrderItemCreateRequest `json:"items"`
}
