package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(router fiber.Router) {
	router.Use(cors.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Api!")
	})

	router.Post("/direction", postDirection)
	router.Get("/direction", getDirections)
}
