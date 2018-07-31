package config

import "github.com/jinzhu/configor"

var Config = struct {
	APPName string `default:"app name"`
	Log struct {
		Level    string `default:"info"`
		Path     string `default:"./logs/"`
		Filename string `default:"producer.log"`
	}
	EnableCoin []string
}{}

func InitConfig(files string) {
	configor.Load(&Config, files)
}
