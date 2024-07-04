package main

import (
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Title string
	Body  string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HEllO")
	})

	app.Get("/post", func(c *fiber.Ctx) error {
		post := Post{
			Title: "hello",
			Body:  "world",
		}
		return c.Status(fiber.StatusOK).JSON(post)
	})
	app.Listen(":3000")
}
