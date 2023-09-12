package models

import "time"

// Note is the note structure that the user will write to the system.
type Note struct {
	Content   string    `json:"content" gorm:"primary_key"`
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tag is a basic structure for tagging a Note.
type Tag struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
