package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func ConnectRedis() {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opts)
}

func GetRDB() *redis.Client {
	return rdb
}
