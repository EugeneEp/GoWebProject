package config

import (
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	cfg := viper.New()

	cfg.SetConfigName("config")
	cfg.SetConfigType("json")
	cfg.AddConfigPath("./")

	cfg.AutomaticEnv()
	if err := cfg.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	return cfg
}
