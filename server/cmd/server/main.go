package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/MartinM2304/hackathon2025/internal/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func gracefulShutdown(app *fiber.App, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
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
			slog.Error("http server error: %s", err.Error())
		}
	}()

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(app, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
