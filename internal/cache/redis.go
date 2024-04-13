package cache

import (
	"github.com/go-redis/redis"
)

// TODO: url from config
func NewRedis() (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr: "redis:6379",
			//Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
