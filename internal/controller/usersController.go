package controller

import (
	"todo/internal/constants"
	"todo/internal/model"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// Hash and salt the password (crypto)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12) // Higher the cost, the more secure, 10 is default but these days that isnt high enough, min should really be 14.

	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, "Couldnt register, please check the request body and try again")
	}

	// Create the user in db
	user := model.User{
		Email:    body.Email,
		Password: string(hash),
	}

	userRes, err := h.storage.CreateUser(&user)
	if err != nil {
		if err.Error() == "That email account is already in use" {
			return utils.JsonResponse(c, "", constants.RESPONSE_EMAIL_CONFLICT)
		}

		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, "Couldnt register, please check the request body and try again")
	}

	// Return with success message
	return utils.JsonResponse(c, userRes, constants.RESPONSE_SUCCESFULLY_CREATED)
}
