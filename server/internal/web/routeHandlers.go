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

	services.RegisterDirectionVote(
		models.Direction{
			Id:     direction.Direction,
			IpAddr: string(c.Request().Host()),
		},
	)

	c.Status(200)
	return nil
}

func postEmoji(c *fiber.Ctx) error {
	emoji := new(models.EmojiJson)

	if err := c.BodyParser(emoji); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterEmojiVote(
		models.Emoji{Id: emoji.Emoji, IpAddr: c.Context().Response.RemoteAddr().String()},
	)

	c.Status(200)
	return nil
}

func getData(c *fiber.Ctx) error {
	err, data := services.GetAggregatedData()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	c.Status(200).JSON(data)
	return nil
}
