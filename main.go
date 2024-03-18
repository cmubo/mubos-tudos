package main

import (
	"log"

	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/initializer"
)

func main() {
	store, err := database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	server := initializer.NewAPIServer(config.Config("PORT"), store)

	server.Start()
}
