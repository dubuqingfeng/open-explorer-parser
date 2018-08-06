package config

import "github.com/jinzhu/configor"

type NodeConfig struct {
	Address  string
	User     string
	Password string
	SSL      bool
	AuthType string
}

type CoinConfig struct {
	Nodes   []NodeConfig
	Network string
}

var Config = struct {
	APPName string `default:"app name"`
	Log struct {
		Level    string `default:"info"`
		Path     string `default:"./logs/"`
		Filename string `default:"producer.log"`
	}
	EnableCoin []string
	ETH        CoinConfig
	ETC        CoinConfig
	XMR        CoinConfig
	BTC        CoinConfig
}{}

func InitConfig(files string) {
	configor.Load(&Config, files)
}
