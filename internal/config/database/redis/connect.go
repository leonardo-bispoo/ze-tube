package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context, url string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	return client
}
