package d7redis

import "github.com/redis/go-redis/v9"

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	}
}
