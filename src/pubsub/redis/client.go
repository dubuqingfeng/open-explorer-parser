package redis

import (
	"github.com/dubuqingfeng/explorer-parser/models/configs"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	Client *redis.Client
)

func Connect(value configs.PubConnConfig) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     value.Address,
		Password: value.Password,
	})
	status := redisClient.Ping()
	if status.Err() != nil {
		panic(status.Err())
	}
	log.WithField("Client", "Redis Client").Info(status)
	Client = redisClient
}

func Publish(value configs.PubConnConfig, prefixKey string, data map[string]string) {
	client := redis.NewClient(&redis.Options{
		Addr:     value.Address,
		Password: value.Password,
		DB:       0,
	})
	for key, message := range data {
		err := client.Publish(prefixKey+key, message).Err()
		if err != nil {
			log.Error(err)
		}
	}
	//defer client.Close()
}
