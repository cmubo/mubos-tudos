package controller

import (
	"time"
	"todo/internal/config"
	"todo/internal/constants"
	"todo/internal/model"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Hash and salt the password (crypto)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12) // Higher the cost, the more secure, 10 is default but these days that isnt high enough, min should really be 14.
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, constants.MSG_REGISTRATION_FAILED_BODY)
	}

	// Create the user in db
	user := model.User{
		Email:    body.Email,
		Password: string(hash),
	}

	userRes, err := h.Storage.CreateUser(&user)
	if err != nil {
		if err.Error() == "That email account is already in use" {
			return utils.JsonResponse(c, "", constants.RESPONSE_EMAIL_CONFLICT)
		}

		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, constants.MSG_REGISTRATION_FAILED_BODY)
	}

	// Return with success message
	return utils.JsonResponse(c, userRes, constants.RESPONSE_SUCCESFULLY_CREATED)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	// Get email and pass from body
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.MSG_LOGIN_FAILED_BODY)
	}

	// Find user
	userRes, err := h.Storage.GetUserFromEmail(body.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, constants.MSG_LOGIN_FAILED_DETAILS)
	}

	// Compare password with saved user pass hash
	if err := bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(body.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, constants.MSG_LOGIN_FAILED_DETAILS)
	}

	// Generate a jwt token
	// Create a new token object, specifying signing method and the claims you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userRes.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, constants.MSG_FAILED_CREATE_TOKEN)
	}

	// send token back in cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = tokenString
	cookie.HTTPOnly = true
	cookie.Secure = false // TODO: if this isnt on local set it to true
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)

	// Set cookie
	c.Cookie(cookie)

	// Return with success message
	return utils.JsonResponse(c, "", constants.RESPONSE_SUCCESFUL_LOGIN)
}

func (h *Handler) Validate(c *fiber.Ctx) error {
	user := c.Locals("user")

	return utils.JsonResponse(c, user, constants.RESPONSE_SUCCESFUL_VALIDATION)
}
