package models

import "time"

type OrderItem struct {
	ID            string    `json:"id" db:"id"`
	Quantity      string    `json:"quantity" db:"quantity" binding:"required,oneof=L M S"`
	Unit_price    float64   `json:"unit_price" db:"unit_price" binding:"required"`
	Created_at    time.Time `json:"created_at" db:"updated_at"`
	Updated_at    time.Time `json:"updated_at" db:"updated_at"`
	Order_item_id string    `json:"order_item_id" db:"order_item_id"`
	Order_id      string    `json:"order_id" db:"order_id"`
}
