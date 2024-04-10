package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

// TODO: url from config
func NewRedis() (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
