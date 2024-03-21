package main

import (
	"log"
	"todo/internal/config"
	"todo/internal/initializer"
)

func main() {
	app := initializer.SetupApi()

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
