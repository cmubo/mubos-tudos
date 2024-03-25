package model

import "time"

type Todo struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title" validate:"required,min=3,max=70"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
