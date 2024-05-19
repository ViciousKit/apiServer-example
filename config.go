package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port      string
	DbUser    string
	Password  string
	Address   string
	DbName    string
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		DbUser:    getEnv("DB_USER", "test"),
		Password:  getEnv("DB+PASSWORD", "secret"),
		Address:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DbName:    getEnv("DB_NAME", "projectManager"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
