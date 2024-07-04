package main

import (
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Title string
	Body  string
}

func main() {
	var post Post
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

	app.Post("/post", func(c *fiber.Ctx) error {
		body := &Post{}
		err := c.BodyParser(body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
		post = Post{
			Title: body.Title,
			Body:  body.Body,
		}
		return c.Status(fiber.StatusOK).JSON(post)
	})
	app.Listen(":3000")
}
