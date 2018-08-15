package config

import (
	"github.com/jinzhu/configor"
	"github.com/dubuqingfeng/explorer-parser/src/models/configs"
)

type CoinConfig struct {
	Network     string
	PublishType string
	PubConn     []configs.PubConnConfig
}

type DB struct {
	Name     string
	User     string `default:"root"`
	Password string `required:"true" env:"DBPassword"`
	Port     uint   `default:"3306"`
}

var Config = struct {
	APPName string `default:"app name"`
	Log struct {
		Level    string `default:"info"`
		Path     string `default:"./logs/"`
		Filename string `default:"consumer.log"`
	}
	EnableCoin []string

	ETH CoinConfig
	ETC CoinConfig
	XMR CoinConfig
	BTC CoinConfig

	Redis struct {
		Address  string
		Password string
	}
}{}

func InitConfig(files string) {
	configor.Load(&Config, files)
}
