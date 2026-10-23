package models

import "time"

type Order struct {
	Order_Date time.Time `json:"order_date" db:"order_date" binding:"required"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	Order_id   string    `json:"order_id" db:"order_id"`
	Table_id   string    `json:"table_id" db:"table_id"`
}
