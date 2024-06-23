package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/auth"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/models"
)

type LoginFields struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserById(c *fiber.Ctx) error {
	var loginFields LoginFields

	if err := c.BodyParser(&loginFields); err != nil {
		fmt.Println("Failed to parse data")
		return err
	}

	var user models.User
	result := db.Storage.Where("email = ?", loginFields.Email).First(&user)
	if result.Error != nil {
		fmt.Println("Failed to get user")
		return result.Error
	}

	userPaylod := auth.UserPayload{FirstName: user.FirstName, Email: user.Email}

	jwtToken, tokenError := auth.GenerateJWT(userPaylod)
	if tokenError != nil {
		return c.Status(500).JSON(tokenError)
	}
	return c.Status(200).JSON(jwtToken)

}
