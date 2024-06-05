package redis

import (
	"app/pkg/config"
	"app/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Client = redis.Client

const Nil = redis.Nil

func NewRedis(config *config.Config, logger *logger.Logger) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	logger.Info("Redis client initialized")
	return rdb
}
