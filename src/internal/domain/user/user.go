// Package user holds the domain entities of the application
package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	AuthIDs  []string  `json:"authIds"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
}

type UserRepository interface {
	GetProfile(token string) (User, error)
}

func InitUser(id uuid.UUID, authIds []string, avatar string, username string) *User {
	return &User{
		ID:       id,
		AuthIDs:  authIds,
		Avatar:   avatar,
		Username: username,
	}
}
