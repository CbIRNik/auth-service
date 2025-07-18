// Package user_test tests the UserHandler
package user_test

import (
	"encoding/json"
	"fmt"
	"auth-service/src/config"
	"auth-service/src/internal/di"
	"auth-service/src/internal/domain/user"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupApp() *fiber.App {
	app := fiber.New()
	envConfig := config.InitEnvConfig()
	db, _ := config.InitPostgres(envConfig)
	deps := di.InitDiContainer(db)
	app.Get("/profile", deps.UserHandler.GetUserProfile)
	return app
}

func TestGetUserProfile_InvalidToken(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	var body map[string]string
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.Equal(t, "Invalid token", body["message"])
}

func TestGetUserProfile_Success(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result user.User
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
	fmt.Println(result)
	assert.NotEqual(t, uuid.Nil, result.ID)
}
