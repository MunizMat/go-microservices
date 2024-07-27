package clients

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

func Init() {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		log.Fatal("REDIS_URL is not defined")
	}

	options, err := redis.ParseURL(redisUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	Redis = redis.NewClient(options)
}
