package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/configs"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/log"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	logger = log.WithLayer("cache")
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		rdb: client,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	rawU, err := json.Marshal(value)
	if err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("error marshaling value: %v", value))
		return err
	}

	if err = c.rdb.Set(ctx, c.formatKey(key), rawU, duration).Err(); err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("error setting value: %v", value))
		return err
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, response interface{}) error {
	val, err := c.rdb.Get(ctx, c.formatKey(key)).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) { // do not log error because of key not found
			logger.Error().Err(err).Msg(fmt.Sprintf("error getting value: %s", key))
		}
		return err
	}

	if err = json.Unmarshal([]byte(val), response); err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("error unmarshalling value: %s", key))
		return err
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	_, err := c.rdb.Del(ctx, c.formatKey(key)).Result()
	if err != nil {
		logger.Error().Err(err).Msg(fmt.Sprintf("error deleting key: %s", key))
		return err
	}
	return nil
}

func (c *Cache) formatKey(key string) string {
	return fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerName, key)
}
