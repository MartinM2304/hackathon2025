package web

import (
	"strconv"

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

func getData(c *fiber.Ctx) error {
	err, direction := services.GetAggregatedDirection()
	if err != nil {
		c.Status(404)
		return err
	}

	c.Status(200).Write([]byte(strconv.Itoa(int(direction))))
	return nil
}
