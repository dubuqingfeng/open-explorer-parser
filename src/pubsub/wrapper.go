package pubsub

import (
	"github.com/dubuqingfeng/explorer-parser/src/models/configs"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/pubsub/redis"
	"github.com/dubuqingfeng/explorer-parser/src/pubsub/rediscluster"
	log "github.com/sirupsen/logrus"
)

const (
	PublishTypeKafka        = "kafka"
	PublishTypeRedis        = "redis"
	PublishTypeRedisCluster = "redis-cluster"
)

type DataWrapper struct {
	prefixKey string
	pubConfig []configs.PubConnConfig
}

func NewDataWrapper(coin string, network string, db []configs.PubConnConfig) *DataWrapper {
	if len(db) == 0 {
		db = config.Config.PubConn
	}
	prefixKey := coin + ":" + network + ":"
	return &DataWrapper{pubConfig: db, prefixKey: prefixKey}
}

func (wrapper *DataWrapper) Publish(data map[string]string) {
	for _, value := range wrapper.pubConfig {
		switch value.PublishType {
		case PublishTypeKafka:
			log.Debug(data)
		case PublishTypeRedis:
			go redis.Publish(value, wrapper.prefixKey, data)
		case PublishTypeRedisCluster:
			rediscluster.Client.Publish("test", "test")
		}
	}
}

func (wrapper *DataWrapper) Subscribe(channel string) {
	for _, value := range wrapper.pubConfig {
		switch value.PublishType {
		case PublishTypeKafka:
			log.Debug(channel)
		case PublishTypeRedis:
			go wrapper.redisSubscribe(channel)
		case PublishTypeRedisCluster:
			log.Debug(channel)
		}
	}
}
func (wrapper *DataWrapper) redisSubscribe(channel string) {
	log.Info(channel)
	//pubsub := redis.Subscribe(channel)
	//msg,_ := pubsub.Receive()
	//fmt.Println("Receive from channel:", msg)
	//var rchannel string
	//var rpayload string
	//
	//for {
	//	msg, err := redis.ReceiveTimeout(time.Second)
	//	if err != nil {
	//		if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) && reflect.TypeOf(err.(*net.OpError).Err).String() == "*net.timeoutError" {
	//			// Timeout, ignore
	//			continue
	//		}
	//		// Actual error
	//		log.Print("Error in ReceiveTimeout()", err)
	//	}
	//
	//	rchannel = ""
	//	rpayload = ""
	//
	//	switch m := msg.(type) {
	//	case *redis.Subscription:
	//		log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
	//		continue
	//	case *redis.Message:
	//		rchannel = m.Channel
	//		rpayload = m.Payload
	//	case *redis.PMessage:
	//		rchannel = m.Channel
	//		rpayload = m.Payload
	//	}
	//
	//	// Process the message
	//	go function(rchannel, rpayload)
	//}
}
