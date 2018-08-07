package pubsub

import (
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/pubsub/rediscluster"
	"github.com/dubuqingfeng/explorer-parser/src/models/configs"
	"github.com/dubuqingfeng/explorer-parser/src/pubsub/redis"
)

type DataWrapper struct {
	prefixKey string
	pubConfig []configs.PubConnConfig
	pubType   string
}

func NewDataWrapper(coin string, network string, pubType string, db []configs.PubConnConfig) *DataWrapper {
	if len(db) == 0 {
		db = config.Config.PubConn
		pubType = config.Config.PublishType
	}
	prefixKey := coin + ":" + network + ":"
	return &DataWrapper{pubConfig: db, pubType: pubType, prefixKey: prefixKey}
}

func (wrapper *DataWrapper) Publish(data map[string]string) {
	switch wrapper.pubType {
	case "kafka":
		log.Debug(data)
	case "redis":
		wrapper.redisPublish(data)
	case "redis-cluster":
		rediscluster.Client.Publish("test", "test")
	}
}

func (wrapper *DataWrapper) redisPublish(data map[string]string) {
	for _, value := range wrapper.pubConfig {
		go func(value configs.PubConnConfig, prefixKey string) {
			redis.Publish(value, prefixKey, data)
		}(value, wrapper.prefixKey)
	}
}
