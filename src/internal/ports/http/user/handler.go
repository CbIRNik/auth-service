// Package user holds the business logic of the application
package user

import (
	"auth-service/src/internal/service/user"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service *user.UserService
}

func (uh *UserHandler) GetUserProfile(c fiber.Ctx) error {
	authToken := c.Get("Authorization")
	user, err := uh.service.GetUserProfile(authToken)
	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
		})
		return nil
	}
	c.Status(200).JSON(fiber.Map{
		"user": user,
	})
	return nil
}

func InitUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
