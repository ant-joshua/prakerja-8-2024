package database

import "github.com/redis/go-redis/v9"

func NewRedis(dbnumber int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",       // no password set
		DB:       dbnumber, // use default DB
	})

	return rdb
}
