package config

import "os"

type Config struct {
	PostgresURL string
	RedisURL    string
}

func LoadEnv() *Config {
	postgresURL, exists := os.LookupEnv("POSTGRES_URL")
	if !exists {
		panic("POSTGRES URL NOT FOUND")
	}

	redisURL, exists := os.LookupEnv("REDIS_URL")
	if !exists {
		panic("REDIS URL NOT FOUND")
	}

	return &Config{
		PostgresURL: postgresURL,
		RedisURL:    redisURL,
	}
}
