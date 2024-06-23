package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/models"
)

func GetPosts(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	var posts []models.Post
	result := db.Storage.Where("user_id = ?", userId).Find(&posts)
	if result.Error != nil {
		return result.Error
	}
	return c.Status(200).JSON(posts)
}
