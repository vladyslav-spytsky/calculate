package conf

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}
