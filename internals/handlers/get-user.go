package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
)

func GetUserById(c *fiber.Ctx) error {
	useridString := c.Params("userId")
	userId, err := strconv.Atoi(useridString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid userid",
		})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"user": db.GetUser(userId),
	})
}
