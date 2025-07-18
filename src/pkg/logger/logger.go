// Package logger holds the business logic of the application
package logger

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} | Time: ${latency}\n",
		TimeFormat: "15:04:05",
		Output:     os.Stdout,
	}))
}
