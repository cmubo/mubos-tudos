package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

func DbErrorSinglularResource(err error) error {
	if err.Error() == "sql: no rows in result set" {
		return fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
	}

	log.Error(err)
	return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
}

func DbErrorMultiResource[T any](err error, emptyResult T) (T, int, error) {
	if err.Error() == "sql: no rows in result set" {
		return emptyResult, 0, nil
	}

	log.Error(err)
	return emptyResult, 0, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
}

func DbGetCount(db *sqlx.DB, tableName string) (int, error) {
	var count int
	// pq package doesnt accept the table name as an argument so we use Sprintf
	// We can only use sprintf here because we control the input. Otherwise the table name would need to be sanitized
	err := db.QueryRowx(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return count, nil
}
