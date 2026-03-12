package cache

import "github.com/redis/go-redis/v9"

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6371",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
