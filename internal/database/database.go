package database

import (
	"fmt"

	"todo/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeDatabase() (*sqlx.DB, error) {
	db, err := startDB()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func startDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s dbname=%s sslmode=%s password=%s",
		config.Config("DB_USER"),
		config.Config("DB_NAME"),
		config.Config("DB_SSLMODE"),
		config.Config("DB_PASSWORD"),
	))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
