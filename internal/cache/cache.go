package cache

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

type cache struct {
	rdb *redis.Client
}

type Cache interface {
	Set(key string, content string) error
	Get(key string) (content string, err error)
	Exist(key string) bool
}

func NewCache(con *redis.Client) Cache {
	return &cache{rdb: con}
}

func (c *cache) Set(key string, content string) error {

	if err := c.rdb.Set(key, content, time.Minute*5); err != nil {
		return err.Err()
	}

	return nil
}

func (c *cache) Get(key string) (content string, err error) {

	if content, err = c.rdb.Get(key).Result(); err != nil {
		if err == redis.Nil {
			return "", errors.New("not exist")
		}
		return "", err
	}

	return content, nil
}

func (c *cache) Exist(key string) bool {

	exists, err := c.rdb.Exists(key).Result()
	if err != nil {
		return true
	}

	return exists > 0
}
