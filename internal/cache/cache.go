package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(con *redis.Client) *Cache {
	return &Cache{rdb: con}
}

func (c *Cache) Set(key string, content string) error {
	ctx := context.Background()

	if err := c.rdb.Set(ctx, key, content, time.Minute*5); err != nil {
		return err.Err()
	}

	return nil
}

func (c *Cache) Get(key string) (content string, err error) {
	ctx := context.Background()

	if content, err = c.rdb.Get(ctx, key).Result(); err != nil {
		if err == redis.Nil {
			return "", errors.New("not exist")
		}
		return "", err
	}

	return content, nil
}

func (c *Cache) Exist(key string) bool {
	ctx := context.Background()

	exists, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		return true
	}

	return exists > 0
}
