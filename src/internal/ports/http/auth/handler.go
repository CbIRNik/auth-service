// Package auth holds the business logic of the application
package auth

import (
	"auth-service/src/internal/service/auth"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	service *auth.AuthService
}

type Body struct {
	Code string `json:"code"`
}

func (ah *AuthHandler) Auth(c fiber.Ctx) error {
	company := c.Params("company")
	body := new(Body)
	if err := c.Bind().Body(body); err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return nil
	}
	authToken, err := ah.service.GetAuthToken(company, body.Code)
	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return nil
	}
	c.Status(200).JSON(fiber.Map{
		"authToken": authToken,
	})
	return nil
}

func InitAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
