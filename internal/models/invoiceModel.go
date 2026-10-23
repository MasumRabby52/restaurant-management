package models

import "time"

type Invoice struct {
	Invoice_id       string    `json:"invoice_id" db:"invoice_id"`
	Order_id         string    `json:"order_id" db:"order_id"`
	PaymentMethod    string    `json:"payment_method" db:"payment_method" binding:"required,oneof=CARD CASH"`
	Payment_status   string    `json:"payment_status" db:"payment_status" binding:"required,oneof=PENDING PAID"`
	Payment_due_date time.Time `json:"payment_due_date" db:"payment_due_date"`
	Created_at       time.Time `json:"created_at" db:"created_at"`
	Updated_at       time.Time `json:"updated_at" db:"updated_at"`
}
