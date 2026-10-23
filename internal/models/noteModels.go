package models

import "time"

type Note struct {
	Text       string    `json:"text" db:"text"`
	Title      string    `json:"title" db:"title"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	Note_id    string    `json:"note_id" db:"note_id"`
}
