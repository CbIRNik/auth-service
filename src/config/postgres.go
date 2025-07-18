// Package config holds the configuration of the application
package config

import (
	"database/sql"
	"fmt"
)

func InitPostgres(envConfig *EnvConfig) (*sql.DB, error) {
	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envConfig.Postgres.Host,
		envConfig.Postgres.Port,
		envConfig.Postgres.User,
		envConfig.Postgres.Password,
		envConfig.Postgres.DBName,
	)
	db, err := sql.Open("postgres", config)
	return db, err
}
