package router

import (
	"todo/internal/controller"
	"todo/internal/middleware"
	"todo/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {

	// Create api group
	api := app.Group("/api")

	api.Get("/", controller.Hello)

	store := storage.NewStorage(db)
	h := controller.NewHandler(store)

	// Users
	api.Post("/register", h.CreateUser)
	api.Post("/login", h.Login)

	// The api/validate route is now authenticated using this
	app.Use("/api/validate", func(c *fiber.Ctx) error {
		return middleware.RequireAuth(c, h)
	})
	api.Get("/validate", h.Validate)

	// Todos
	api.Get("/todo", h.GetTodos)
	api.Post("/todo", h.CreateTodo)
	api.Put("/todo", h.UpdateTodo)
	api.Get("/todo/:id", h.GetTodo)
	api.Delete("/todo/:id", h.DeleteTodo)
}
