package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedis(dbnumber int) *redis.Client {

	redisAddr := fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: viper.GetString("REDIS_PASS"), // no password set
		DB:       dbnumber,                      // use default DB
	})

	return rdb
}
