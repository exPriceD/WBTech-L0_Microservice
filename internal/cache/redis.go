package cache

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Url,
		Password: cfg.Redis.Password,
		DB:       int(cfg.Redis.DB),
	})

	return redisClient
}
