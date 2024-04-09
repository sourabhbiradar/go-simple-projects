package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

var Appconfig *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Error on parsing Config file", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
