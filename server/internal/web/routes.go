package web

import (
	"github.com/MartinM2304/hackathon2025/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sync"
)

// Thread-safe list of directions
var (
	directions []models.DirectionJson
	mutex      sync.Mutex // Prevents race conditions
)

func Register(router fiber.Router) {
	router.Use(cors.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Api!")
	})

	router.Post("/direction", func(c *fiber.Ctx) error {
		direction := new(models.DirectionJson)

		if err := c.BodyParser(direction); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		mutex.Lock()
		directions = append(directions, *direction)
		mutex.Unlock()

		return c.Status(200).JSON(fiber.Map{
			"message":   "Direction received successfully",
			"direction": direction,
		})
	})

	router.Get("/direction", func(c *fiber.Ctx) error {
		mutex.Lock()
		defer mutex.Unlock()

		return c.JSON(directions)
	})
}
