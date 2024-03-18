package controller

import "todo/internal/storage"

type Handler struct {
	storage storage.Storage
}

func NewHandler(s storage.Storage) *Handler {
	return &Handler{
		storage: s,
	}
}
