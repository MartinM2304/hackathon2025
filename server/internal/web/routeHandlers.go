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
		models.Emoji{Id: emoji.Emoji, IpAddr: string(c.Request().Host())},
	)

	c.Status(200)
	return nil
}

func postSound(c *fiber.Ctx) error {
	sound := new(models.SoundJson)

	if err := c.BodyParser(sound); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterSoundVote(
		models.Sound{Id: sound.Sound, IpAddr: string(c.Request().Host())},
	)

	c.Status(200)
	return nil
}

func getData(c *fiber.Ctx) error {
	data := services.GetAggregatedData()

	err := c.Status(200).JSON(data)
	if err != nil {
		c.Status(500)
		return err
	}
	return nil
}

func getStats(c *fiber.Ctx) error {
	entropy, err := services.CalculateEnthropy()
	if err != nil {
		c.Status(500)
		return err
	}

	body := models.StatsJson{
		Enthropy: entropy,
	}

	err = c.Status(200).JSON(body)
	if err != nil {
		c.Status(500)
		return err
	}

	return nil
}
