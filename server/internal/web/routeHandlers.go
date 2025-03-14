package web

import (
	"github.com/MartinM2304/hackathon2025/internal/models"
	"github.com/gofiber/fiber/v2"
	"sync"
)

// Thread-safe list of directions
var (
	directions []models.DirectionJson
	mutex      sync.Mutex // Prevents race conditions
)

func postDirection(c *fiber.Ctx) error {
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
}

func getDirections(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	return c.JSON(directions)
}
