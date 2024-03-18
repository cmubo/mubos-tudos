package initializer

import (
	"log"
	"todo/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	listenAddr string
	db         *sqlx.DB
}

func NewAPIServer(listenAddr string, db *sqlx.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Start() {
	app := fiber.New()

	// Put top level middleware here

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.SetupRoutes(app, s.db)

	log.Fatal(app.Listen(":3000"))
}
