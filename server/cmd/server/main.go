package main

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
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

	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		log.Printf("Connection event 1 - UUID: %s\n", ep.Kws.GetUUID())
	})

	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		log.Printf("Disconnection event - UUID: %s\n", ep.Kws.GetUUID())
	})

	router := app.Group("/api")
	web.Register(router)

	wsRouter := app.Group("/socket.io")
	wsRouter.Use(func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	wsRouter.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	wsRouter.Get("/", socketio.New(func(kws *socketio.Websocket) {}))

	done := make(chan bool, 1)

	go func() {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			slog.Warn("Falling back to port 3000")
			port = 3000
		}

		err = app.Listen(fmt.Sprintf(":%d", port))
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
				services.Aggregate()
			case m := <-services.NotificationChannel:
				socketio.Broadcast([]byte(m), socketio.TextMessage)
			}
		}
	}()

	go gracefulShutdown(app, done)

	<-done
	log.Println("Graceful shutdown complete.")
}
