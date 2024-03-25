package model

import "time"

type User struct {
	Id        uint      `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type CreateUser struct {
	Email    string
	Password string
}

type LoginUser struct {
	Email    string
	Password string
}
