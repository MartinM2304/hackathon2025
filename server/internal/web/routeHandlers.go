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

	c.Status(200)
	return nil
}

func postEmoji(c *fiber.Ctx) error {
	emoji := new(models.EmojiJson)

	if err := c.BodyParser(emoji); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterEmojiVote(emoji.Emoji)

	c.Status(200)
	return nil
}

func getData(c *fiber.Ctx) error {
	data := services.GetAggregatedData()

	c.Status(200).JSON(data)
	return nil
}
