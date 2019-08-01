package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var RC *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	RC = client
}

func GetClient() *redis.Client {
	return RC
}
