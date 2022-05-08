package cachedb

import (
	"github.com/go-redis/redis/v8"
	"github.com/husdeli/DrawPadService.git/internal/config"
	redisClient "github.com/husdeli/DrawPadService.git/pkg/clients/redis"
)

func NewCacheDb() *redis.Client {
	cfg := config.GetConfig()

	return redisClient.NewClient(cfg.CacheDbHost, cfg.CacheDbPort, "")
}
