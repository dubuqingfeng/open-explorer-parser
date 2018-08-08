package config

import "github.com/jinzhu/configor"

var Config = struct {
	APPName string `default:"app name"`
	Log struct {
		Level    string `default:"info"`
		Path     string `default:"./logs/"`
		Filename string `default:"consumer.log"`
	}
	EnableCoin []string

	Redis struct {
		Address  string
		Password string
	}

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}{}

func InitConfig(files string) {
	configor.Load(&Config, files)
}
