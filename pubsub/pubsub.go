package pubsub

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitPubSub() {
	redisClient = redis.NewClient(&redis.Options{
		// TODO: .env
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func ensureRedisInitialized() {
	if redisClient == nil {
		panic("Redis Client not initialized. Call InitPubSub() on initialization.")
	}
}

func Pub(key string, message []byte) {
	ensureRedisInitialized()
	err := redisClient.Publish(context.Background(), key, message).Err()
	if err != nil {
		fmt.Printf("[pubsub] Error on Pub(): %v\n", err)
	}
}

func Sub(key string) *redis.PubSub {
	ensureRedisInitialized()
	subscriber := redisClient.Subscribe(context.Background(), key)
	return subscriber
}
