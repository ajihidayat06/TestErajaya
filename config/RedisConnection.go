package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var Client *redis.Client

func StartConnectingRedis() {
	fmt.Println("Start Connecting to Redis...")
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(Client.Context()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis!")
}
