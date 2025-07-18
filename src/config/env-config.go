package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	Server   Server
	Postgres Postgres
	Redis    Redis
	Github   Github
	Google   Google
	JWT      JWT
}

type Server struct {
	Host string
	Port string
}
type Postgres struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
}
type Redis struct {
	Host string
	Port string
}
type Github struct {
	ClientID     string
	ClientSecret string
}
type Google struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}
type JWT struct {
	JWTSecret string
}

func InitEnvConfig() *EnvConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	envConfig := EnvConfig{
		Server{
			Host: viper.GetString("HOST"),
			Port: viper.GetString("PORT"),
		},
		Postgres{
			DBName:   viper.GetString("POSTGRES_DB"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
		},
		Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetString("REDIS_PORT"),
		},
		Github{
			ClientID:     viper.GetString("GITHUB_CLIENT_ID"),
			ClientSecret: viper.GetString("GITHUB_CLIENT_SECRET"),
		},
		Google{
			ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
			ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
			CallbackURL:  viper.GetString("GOOGLE_CALLBACK_URL"),
		},
		JWT{
			JWTSecret: viper.GetString("JWT_SECRET"),
		},
	}
	return &envConfig
}
