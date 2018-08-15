package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/models/configs"
)

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

//func Publish(value configs.PubConnConfig, prefixKey string, data map[string]string) {
//	client := redis.NewClient(&redis.Options{
//		Addr:     value.Address,
//		Password: value.Password,
//		DB:       0,
//	})
//	for key, message := range data {
//		err := client.Publish(prefixKey+key, message).Err()
//		if err != nil {
//			log.Error(err)
//		}
//	}
//	//defer client.Close()
//}