package pubsub

import (
	"context"
	"errors"
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

func ensureRedisInitialized() error {
	if redisClient == nil {
		return errors.New("redis Client not initialized, need to call InitPubSub() on initialization")
	}
	return nil
}

func Pub(key string, message []byte) error {
	err := ensureRedisInitialized()
	if err != nil {
		return err
	}

	err = redisClient.Publish(context.Background(), key, message).Err()
	if err != nil {
		return fmt.Errorf("[pubsub] Error on Pub(): %v", err)
	}
	return nil
}

func Sub(key string) (*redis.PubSub, error) {
	err := ensureRedisInitialized()
	if err != nil {
		return nil, err
	}

	subscriber := redisClient.Subscribe(context.Background(), key)
	return subscriber, nil
}
