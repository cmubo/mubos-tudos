package router

import (
	"todo/internal/controller"
	"todo/internal/middleware"
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
	authenticatedRoute := func(c *fiber.Ctx) error {
		// Checks whether there is a valid JWT token sent and it hasnt expired.
		return middleware.RequireAuth(c, h)
	}

	// Users
	api.Post("/register", h.CreateUser)
	api.Post("/login", h.Login)

	app.Use("/api/validate", authenticatedRoute)
	api.Get("/validate", h.Validate)

	// Todos
	api.Use("/todo", authenticatedRoute)
	api.Get("/todo", h.GetTodos)
	api.Post("/todo", h.CreateTodo)
	api.Put("/todo", h.UpdateTodo)
	api.Get("/todo/:id", h.GetTodo)
	api.Delete("/todo/:id", h.DeleteTodo)
}
