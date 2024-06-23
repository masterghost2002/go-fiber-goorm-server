package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/db"
	"github.com/masterghost2002/go-todo/internals/handlers"
)

func main() {

	err := db.StorageInit()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// group for auth request
	auth := app.Group("/auth")
	auth.Post("/register", handlers.RegisterUser)
	auth.Get("/:userId", handlers.GetUserById)

	app.Listen(":3000")
}
