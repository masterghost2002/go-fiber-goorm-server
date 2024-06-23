package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/go-todo/internals/auth"
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
	authRouter := app.Group("/auth")
	authRouter.Post("/register", handlers.RegisterUser)
	authRouter.Post("/login", handlers.GetUserById)

	// group for post request
	post := app.Group("/post", auth.AuthUser)
	post.Get("/", handlers.GetPosts)
	post.Post("/add-post", handlers.AddPost)

	app.Listen(":3000")
}
