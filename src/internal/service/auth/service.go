// Package auth holds the business logic of the application
package auth

import (
	"errors"
	"fmt"
	"auth-service/src/internal/domain/user"
)

type AuthService struct {
	repository user.UserRepository
	strategies map[string]AuthStrategy
}

type AuthStrategy interface {
	Auth(code string) (string, error)
}

type GithubAuthStrategy struct{}

func (g *GithubAuthStrategy) Auth(code string) (string, error) {
	return "github", nil
}

type GoogleAuthStrategy struct{}

func (g *GoogleAuthStrategy) Auth(code string) (string, error) {
	return "google", nil
}

func (as *AuthService) register(name string, strategy AuthStrategy) {
	as.strategies[name] = strategy
}

func (as *AuthService) getStrategy(name string) (AuthStrategy, error) {
	strategy, ok := as.strategies[name]
	if !ok {
		fmt.Println("strategy not found")
		return nil, errors.New("strategy not found")
	}
	return strategy, nil
}

func (as *AuthService) GetAuthToken(company string, code string) (string, error) {
	strategy, err := as.getStrategy(company)
	if err != nil {
		return "", err
	}
	return strategy.Auth(code)
}

func InitAuthService(repository user.UserRepository) *AuthService {
	authService := AuthService{
		repository: repository,
		strategies: make(map[string]AuthStrategy),
	}
	authService.register("github", &GithubAuthStrategy{})
	authService.register("google", &GoogleAuthStrategy{})
	return &authService
}
