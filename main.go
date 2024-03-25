package main

import (
	"log"
	"todo/internal/config"
	"todo/internal/initializer"
)

// @title		Mubos Todos API
// @version		1.0
// @description	Just a simple API made for a todo application
func main() {
	app := initializer.SetupApi()

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
