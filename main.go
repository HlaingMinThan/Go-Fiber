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

	postApis := app.Group("/post")

	postApis.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(post)
	})

	postApis.Post("/", func(c *fiber.Ctx) error {
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
