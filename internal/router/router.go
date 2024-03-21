package router

import (
	"todo/internal/controller"
	"todo/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {

	app.Get("/hello", controller.Hello)

	// Create api group
	api := app.Group("/api")

	store := storage.NewStorage(db)
	h := controller.NewHandler(store)

	// Todos
	api.Get("/todo", h.GetTodos)
	api.Post("/todo", h.CreateTodo)
	api.Put("/todo", h.UpdateTodo)
	api.Get("/todo/:id", h.GetTodo)
	api.Delete("/todo/:id", h.DeleteTodo)
}
