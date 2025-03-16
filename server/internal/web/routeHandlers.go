package web

import (
	"github.com/MartinM2304/hackathon2025/internal/models"
	"github.com/MartinM2304/hackathon2025/internal/services"
	"github.com/gofiber/fiber/v2"
)

func postDirection(c *fiber.Ctx) error {
	body := new(models.DirectionJson)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterDirectionVote(body.Direction)

	c.Status(200)
	return nil
}

func postEmoji(c *fiber.Ctx) error {
	body := new(models.EmojiJson)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterEmojiVote(body.Emoji)

	c.Status(200)
	return nil
}

func postSound(c *fiber.Ctx) error {
	body := new(models.SoundJson)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	services.RegisterSoundVote(body.Sound)

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

	emojiData, err := services.CalculateStats("emoji")
	if err != nil {
		c.Status(500)
		return err
	}
	directionData, err := services.CalculateStats("direction")
	if err != nil {
		c.Status(500)
		return err
	}
	soundData, err := services.CalculateStats("sound")
	if err != nil {
		c.Status(500)
		return err
	}

	body := models.StatsJson{
		Enthropy:  entropy,
		Emoji:     emojiData,
		Direction: directionData,
		Sound:     soundData,
	}

	err = c.Status(200).JSON(body)
	if err != nil {
		c.Status(500)
		return err
	}

	return nil
}
