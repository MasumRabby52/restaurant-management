package models

import "time"

type User struct {
	First_name string `json:"first_name" db:"first_name" binding:"required"`
	Last_name  string `json:"last_name" db:"last_name" binding:"required"`
	Password   string `json:"-" db:"password" binding:"required"`
	Email      string `json:"email" db:"email" binding:"required"`
	Phone      string `json:"phone" db:"phone"`
	// Token         string    `json:"token" db:"token"`
	// Refresh_Token string    `json:"refresh_token" db:"refresh_token"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	User_id    string    `json:"user_id" db:"user_id"`
}
