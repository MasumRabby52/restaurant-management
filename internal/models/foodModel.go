package models

import "time"

type Food struct {
	Name       string    `json:"name" db:"name"`
	Price      float64   `json:"price" db:"price"`
	Food_image string    `json:"food_price" db:"food_price"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	Food_id    string    `json:"food_id" db:"food_id"`
	Menu_id    string    `json:"menu_id" db:"menu_id"`
}
