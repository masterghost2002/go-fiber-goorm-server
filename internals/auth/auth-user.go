package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/models"
)

func AuthUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	// Check if the header is present
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header missing",
		})
	}

	// Check if the header starts with "Bearer "
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	token := authHeader[len(bearerPrefix):]

	user, err := ValidateJWT(token)
	if err != nil {
		return err
	}
	var userFromDb models.User
	result := db.Storage.Where("email = ?", user.Email).First(&userFromDb)
	if result.Error != nil {
		return result.Error
	}
	c.Locals("user_email", userFromDb.Email)
	c.Locals("user_id", userFromDb.ID)
	return c.Next()
}
