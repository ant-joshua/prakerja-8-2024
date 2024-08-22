package models

type Product struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Category   Category `json:"category"` // embeded struct
	CategoryID int      `json:"category_id"`
	Price      int      `json:"price"`
	Stock      int      `json:"stock"`
}
