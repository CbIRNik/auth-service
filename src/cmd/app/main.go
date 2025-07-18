package main

import (
	"fmt"
	"auth-service/src/config"
	"auth-service/src/internal/di"
	router "auth-service/src/internal/ports/http"
	"auth-service/src/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	logger.SetupLogger(app)
	envConfig := config.InitEnvConfig()
	db, _ := config.InitPostgres(envConfig)
	deps := di.InitDiContainer(db)
	router.InitRouter(app, deps)
	addr := fmt.Sprintf("%s:%s", envConfig.Server.Host, envConfig.Server.Port)
	app.Listen(addr, fiber.ListenConfig{
		EnablePrefork: true,
	})
}
