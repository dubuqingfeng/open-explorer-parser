package rediscluster

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	Client *redis.ClusterClient
)

func Connect(db string, address []string) {
	fmt.Println(address)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
	})
	status := client.Ping()
	if status.Err() != nil {
		panic(status.Err())
	}
	log.WithField("Client", "Redis Cluster Client").Info(status)
	Client = client
}
