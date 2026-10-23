package models

import "time"

type Table struct {
	Number_of_guests int       `json:"number_of_guests" db:"number_of_guests"`
	Table_number     int       `json:"table_number" db:"table_number"`
	Created_at       time.Time `json:"created_at" db:"created_at"`
	Updated_at       time.Time `json:"updated_at" db:"updated_at"`
	Table_id         string    `json:"table_id" db:"table_id"`
}
