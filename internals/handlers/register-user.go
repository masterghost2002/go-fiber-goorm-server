package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/models"
)

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	var userData UserData
	if err := c.BodyParser(&userData); err != nil {
		return err
	}

	user := models.User{FirstName: userData.FirstName, LastName: &userData.LastName, Email: userData.Email, Password: userData.Password}

	result := db.Storage.Create(&user)
	if result.Error != nil {
		return c.Status(424).JSON(fiber.Map{
			"error": result.Error,
		})
	}
	return c.SendStatus(200)
}
