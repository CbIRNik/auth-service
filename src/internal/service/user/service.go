// Package user holds the business logic of the application
package user

import (
	"auth-service/src/internal/domain/user"
)

type UserService struct {
	repository user.UserRepository
}

func (us *UserService) GetUserProfile(token string) (user.User, error) {
	result, err := us.repository.GetProfile(token)
	return result, err
}

func InitUserService(repository user.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}
