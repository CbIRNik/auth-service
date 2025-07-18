// Package user holds the business logic of the application
package user

import (
	"database/sql"
	"auth-service/src/internal/domain/user"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) GetProfile(token string) (user.User, error) {
	return *user.InitUser(uuid.New(), []string{"123", "345"}, "123", "123"), nil
}

func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
