package controller

import (
	"todo/internal/storage"
	"todo/internal/utils"
)

type Handler struct {
	storage   storage.Storage
	validator *utils.Validator
}

func NewHandler(s storage.Storage) *Handler {
	v := utils.NewValidator()
	return &Handler{
		storage:   s,
		validator: v,
	}
}
