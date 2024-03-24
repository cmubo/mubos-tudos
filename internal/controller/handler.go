package controller

import (
	"todo/internal/storage"
	"todo/internal/utils"
)

type Handler struct {
	Storage   storage.Storage
	Validator *utils.Validator
}

func NewHandler(s storage.Storage) *Handler {
	v := utils.NewValidator()
	return &Handler{
		Storage:   s,
		Validator: v,
	}
}
