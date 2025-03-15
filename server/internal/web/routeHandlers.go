package web

import (
	"github.com/MartinM2304/hackathon2025/internal/models"
	"github.com/MartinM2304/hackathon2025/internal/services"
	"github.com/gofiber/fiber/v2"
)

func postDirection(c *fiber.Ctx) error {
	direction := new(models.DirectionJson)

	if err := c.BodyParser(direction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterDirectionVote(direction.Direction)

	return c.Status(200).JSON(fiber.Map{
		"message":   "Direction received successfully",
		"direction": direction,
	})
}

func getDirections(c *fiber.Ctx) error {
	directions := services.GetDirections()
	return c.JSON(directions)
}
