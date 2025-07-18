// Package di holds the dependency injection of the application
package di

import (
	"database/sql"
	userhandler "auth-service/src/internal/ports/http/user"
	authhandler "auth-service/src/internal/ports/http/auth"
	userrepository "auth-service/src/internal/repository/user"
	authservice "auth-service/src/internal/service/auth"
	userservice "auth-service/src/internal/service/user"
)

type DiContainer struct {
	UserHandler *userhandler.UserHandler
	AuthHandler *authhandler.AuthHandler
}

func InitUserDeps(db *sql.DB) *userhandler.UserHandler {
	userRepository := userrepository.InitUserRepository(db)
	userService := userservice.InitUserService(userRepository)
	return userhandler.InitUserHandler(userService)
}

func InitAuthDeps(db *sql.DB) *authhandler.AuthHandler {
	userRepository := userrepository.InitUserRepository(db)
	authService := authservice.InitAuthService(userRepository)
	return authhandler.InitAuthHandler(authService)
}

func InitDiContainer(db *sql.DB) *DiContainer {
	return &DiContainer{
		InitUserDeps(db),
		InitAuthDeps(db),
	}
}
