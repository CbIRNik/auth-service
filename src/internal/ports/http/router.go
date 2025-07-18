// Package router holds the business logic of the application
package router

import (
	"auth-service/src/internal/di"

	"github.com/gofiber/fiber/v3"
)

func InitRouter(app *fiber.App, deps *di.DiContainer) {
	userGroup := app.Group("/user")
	userGroup.Get("/profile", deps.UserHandler.GetUserProfile)
	userGroup.Post("/auth/:company", deps.AuthHandler.Auth)
}
