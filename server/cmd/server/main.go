package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/MartinM2304/hackathon2025/internal/services"
	"github.com/MartinM2304/hackathon2025/internal/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func gracefulShutdown(app *fiber.App, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006",
		TimeZone:   "Europe/Sofia",
	}))

	router := app.Group("/api")
	web.Register(router)

	done := make(chan bool, 1)

	go func() {
		err := app.Listen(":3000")
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err.Error()))
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := services.Aggregate()
				if err != nil {
					fmt.Println("ERRRRRRRROOOOOOORRRRRR")
				}
			}
		}
	}()

	go gracefulShutdown(app, done)

	<-done
	log.Println("Graceful shutdown complete.")
}
