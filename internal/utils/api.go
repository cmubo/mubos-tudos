package utils

import (
	"todo/internal/types"

	"github.com/gofiber/fiber/v2"
)

func JsonResponse[T any](ctx *fiber.Ctx, data T, rData types.Response) error {
	return ctx.JSON(fiber.Map{
		"status":  rData.Status,
		"message": rData.Message,
		"data":    data,
	})
}
