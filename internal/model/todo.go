package model

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=3,max=70"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
