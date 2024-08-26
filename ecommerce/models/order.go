package models

type Order struct {
	ID     int  `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"foreignKey:UserID" json:"user_id"`
}

type OrderItems struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"foreignKey:OrderID" json:"order_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	ProductID int     `gorm:"foreignKey:ProductID" json:"product_id"`
	Qty       int     `gorm:"qty" json:"qty"`
	Price     int     `gorm:"price" json:"price"`
}
