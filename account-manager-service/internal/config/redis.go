package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedisClient(viperConfig *viper.Viper) *redis.Client {
	address := viperConfig.GetString("REDIS_ADDRESS")
	port := viperConfig.GetString("REDIS_PORT")
	db := viperConfig.GetInt("REDIS_DB")
	password := viperConfig.GetString("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", address, port),
		DB:       db,
		Password: password,
	})

	// Delete all redis value when app restarting
	client.FlushDB(context.Background())

	return client
}
