package models

import "time"

type Menu struct {
	Name       string    `json:"name" db:"name"`
	Category   string    `json:"category" db:"category" binding:"required"`
	Start_Date time.Time `json:"start_date" db:"start_date"`
	End_Date   time.Time `json:"end_date" db:"end_date"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	Menu_id    string    `json:"menu_id" db:"menu_id"`
}
