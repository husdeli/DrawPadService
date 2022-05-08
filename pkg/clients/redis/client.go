package redisClient

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewClient(host string, port int, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})
}
