package main

import (
	"log/slog"

	"github.com/MartinM2304/hackathon2025/internal/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006",
		TimeZone:   "Europe/Sofia",
	}))

	router := app.Group("/api")
	web.Register(router)

	err := app.Listen(":3000")
	if err != nil {
		slog.Error("Error starting the server %s", err.Error())
	}
}
