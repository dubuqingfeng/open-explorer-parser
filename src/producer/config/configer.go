package config

import (
	"github.com/dubuqingfeng/explorer-parser/models/configs"
	"github.com/jinzhu/configor"
)

type NodeConfig struct {
	Address  string
	User     string
	Password string
	SSL      bool
	AuthType string
	Type     string
}

type CoinConfig struct {
	Nodes   []NodeConfig
	Network string
	PubConn []configs.PubConnConfig
}

var Config = struct {
	APPName string `default:"app name"`
	Log     struct {
		Level    string `default:"info"`
		Path     string `default:"./logs/"`
		Filename string `default:"producer.log"`
	}
	PubConn    []configs.PubConnConfig
	EnableCoin []string
	ETH        CoinConfig
	ETC        CoinConfig
	XMR        CoinConfig
	BTC        CoinConfig
}{}

func InitConfig(files string) {
	if err := configor.Load(&Config, files); err != nil {
		panic(err)
	}
}
