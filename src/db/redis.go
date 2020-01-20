package db

import (
	"github.com/dubuqingfeng/explorer-parser/consumer/config"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	Client *redis.Client
)

func Connect(db string, address []string) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
	})
	status := redisClient.Ping()
	if status.Err() != nil {
		panic(status.Err())
	}
	log.WithField("Client", "Redis Client").Info(status)
	Client = redisClient
}
