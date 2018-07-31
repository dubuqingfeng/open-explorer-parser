package config

import "github.com/jinzhu/configor"

var Config = struct {
	APPName string `default:"app name"`

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

func InitConfig() {
	configor.Load(&Config, "./config/producer.json")
}
