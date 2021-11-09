package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/renanpalmeira/me-lembra-bot/internal"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	// Automatic load environment variables from .env.
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Set global zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	_ = zap.ReplaceGlobals(logger)

	reminder := internal.NewReminder()

	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./frontend/build")
	app.Post("/reminder-me", func(c *fiber.Ctx) error {
		var payload internal.ReminderPayload
		if err := c.BodyParser(&payload); err != nil {
			zap.L().Error("reminder-me-endpoint", zap.Error(err), zap.ByteString("payload", c.Body()))
			return c.SendStatus(http.StatusBadRequest)
		}

		payload.CreatedAt = time.Now().UTC()

		err := reminder.Add(payload)
		if err != nil {
			zap.L().Error("reminder-me-reminder-sms", zap.Error(err), zap.Reflect("payload", payload))
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.SendStatus(http.StatusOK)
	})

	go func() {
		zap.L().Info("cron", zap.String("description", "up and running cron"))

		c := cron.New()
		_, _ = c.AddFunc("@every 30s", func() { reminder.Run() })
		c.Start()
	}()

	if err := app.Listen(":" + os.Getenv("PORT")); !errors.Is(err, http.ErrServerClosed) {
		zap.L().Error("http-server", zap.Error(err))
	}
}
