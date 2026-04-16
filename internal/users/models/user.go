package models

import "time"

type User struct {
	ID        int64     `json:"ID"`
	Name      string    `json:"name" binding:"required,min=5,max=200"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
