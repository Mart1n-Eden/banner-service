package cache

import (
	"banner-service/internal/config"
	"github.com/go-redis/redis"
)

func NewRedis(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.URL,
			Password: "",
			DB:       0,
		})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
