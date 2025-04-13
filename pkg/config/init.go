package config

import (
	"github.com/spf13/viper"

	"github.com/woxQAQ/upload-server/pkg/types"
)

var cfg *types.AppConfig

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *types.AppConfig {
	return cfg
}
