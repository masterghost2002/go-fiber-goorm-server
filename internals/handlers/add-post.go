package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/models"
)

type PostData struct {
	Title   string `j{son:"title"`
	Content string `json:"content"`
}

func AddPost(c *fiber.Ctx) error {
	var postData PostData
	var userIdU uint = c.Locals("user_id").(uint)
	var userId int = int(userIdU)
	if err := c.BodyParser(&postData); err != nil {
		return err
	}

	post := models.Post{Title: postData.Title, Content: postData.Content, UserId: userId}

	result := db.Storage.Create(&post)
	if result.Error != nil {
		return c.Status(424).JSON(fiber.Map{
			"error": result.Error,
		})
	}
	return c.Status(200).JSON(post)
}
