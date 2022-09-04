package internal

import (
	"github.com/go-redis/redis/v8"
)

func NewRedis(url string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url, // 
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
