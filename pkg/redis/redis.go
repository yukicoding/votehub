package redis

import (
	"fmt"
	"strconv"
	config "yukicoding/voteHub/configs"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Init(config config.RedisConfig) error {
	client = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       0, // Using default DB
	})

	// This block is for selecting a specific Redis database if provided in the configuration
	if config.RedisDb != "" {
		// Convert the RedisDb string to an integer
		dbNumber, err := strconv.Atoi(config.RedisDb)
		if err != nil {
			return fmt.Errorf("invalid Redis DB number: %w", err)
		}
		// Use the SELECT command to switch to the specified database
		if err := client.Do(client.Context(), "SELECT", dbNumber).Err(); err != nil {
			return fmt.Errorf("failed to select Redis DB: %w", err)
		}
	}

	// 可以在这里添加一个 Ping 操作来确保连接成功
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return nil
}

func GetClient() *redis.Client {
	return client
}
