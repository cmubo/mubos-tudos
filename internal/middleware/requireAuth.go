package middleware

import (
	"fmt"
	"time"
	"todo/internal/config"
	"todo/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *fiber.Ctx, h *controller.Handler) error {
	// Get the JWT cookie
	authCookie := c.Cookies("authorization")
	if authCookie == "" {
		log.Error("No auth cookie")
		return fiber.NewError(fiber.StatusUnauthorized, "No authorization cookie")
	}

	// Decode and validated the cookie
	token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Config("JWT_SECRET")), nil
	})

	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusUnauthorized, "Authentication failed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiration := claims["exp"]
		userId, ok := claims["sub"].(float64)
		if !ok {
			log.Error("sub is not a float64")
			return fiber.NewError(fiber.StatusUnauthorized, "Authentication failed")
		}

		// Check the expiration
		if float64(time.Now().Unix()) > expiration.(float64) {
			return fiber.NewError(fiber.StatusUnauthorized, "Authentication failed")
		}

		// FInd user with token subject
		user, err := h.Storage.GetUser(int(userId))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Authentication failed")
		}

		// Attach to the request
		c.Locals("user", user)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Authentication failed")
	}

	// Continue
	return c.Next()
}
