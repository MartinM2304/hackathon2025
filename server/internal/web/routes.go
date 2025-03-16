package web

import (
	"github.com/MartinM2304/hackathon2025/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(router fiber.Router) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "https://client.g8row.xyz, http://localhost:3001",    // Allow all origins (or specify your frontend URL)
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",                        // Allowed methods
		AllowHeaders:     "Content-Type,Authorization,Accept,X-Requested-With", // Allowed headers
		ExposeHeaders:    "Authorization",                                      // Exposed headers
		AllowCredentials: true,                                                 // Allow credentials (if needed)
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Api!")
	})

	router.Post("/direction", postDirection)
	router.Post("/emoji", postEmoji)
	router.Post("/sound", postSound)

	router.Get("/data", getData)
	router.Get("/stats", getStats)
	router.Get("/dump", func(c *fiber.Ctx) error {
		votes, err := database.DumpVotesTable()
		if err != nil {
			c.Status(500)
			return err
		}

		err = c.Status(200).JSON(votes)
		if err != nil {
			c.Status(500)
			return err
		}

		return nil
	})
}
