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

// CreateUser godoc
//
//	@Summary		Register an account
//	@Description	Register a new user account
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body	model.CreateUser	true	"Add user"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/register [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.MSG_REGISTRATION_FAILED_BODY)
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
			return fiber.NewError(fiber.StatusBadRequest, constants.MSG_EMAIL_ADDRESS_IN_USE)
		}

		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, constants.MSG_REGISTRATION_FAILED_BODY)
	}

	// Return with success message
	return utils.JsonResponse(c, userRes, constants.RESPONSE_SUCCESFULLY_CREATED)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login to your user account
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body	model.LoginUser		true	"Login user"
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/login [post]
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

	jwtToken, err := createJwtToken(int(userRes.Id))
	if err != nil {
		return err
	}

	setAuthCookie(c, jwtToken)

	// Return with success message
	return utils.JsonResponse(c, "", constants.RESPONSE_SUCCESFUL_LOGIN)
}

// Login godoc
//
//	@Summary		Validate
//	@Description	Validate that user account is still logged in and refresh JWT token. This will require the authorization cookie as part of the request to be successful.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/validate [post]
func (h *Handler) Validate(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	// Renew the JWT token
	jwtToken, err := createJwtToken(int(user.Id))
	if err != nil {
		return err
	}

	setAuthCookie(c, jwtToken)

	return utils.JsonResponse(c, user, constants.RESPONSE_SUCCESFUL_VALIDATION)
}

func createJwtToken(userId int) (string, error) {
	// Create a new token object, specifying signing method and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 14).Unix(), // 14 days
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, constants.MSG_FAILED_CREATE_TOKEN)
	}

	return tokenString, nil
}

func setAuthCookie(c *fiber.Ctx, token string) {
	// send token cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Secure = false                                // TODO: environment
	cookie.Expires = time.Now().Add(time.Hour * 24 * 14) // 14 days

	// Set cookie
	c.Cookie(cookie)
}
