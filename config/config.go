package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	port     string
	esConfig ElasticSearchConfig
}

func Load() Config {
	viper.AutomaticEnv()

	return Config{
		port:     getStringValue("APP_PORT"),
		esConfig: newESConfig(),
	}
}

func (conf Config) GetPortNumber() string {
	return conf.port
}
