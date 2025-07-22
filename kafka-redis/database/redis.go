package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected.")
}
