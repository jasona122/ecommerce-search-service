package config

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/spf13/viper"
)

type Config struct {
	port          string
	hystrixConfig hystrix.CommandConfig
	esConfig      ElasticSearchConfig
}

func Load() Config {
	viper.AutomaticEnv()

	return Config{
		port:          getStringValue("APP_PORT"),
		hystrixConfig: newHystrixConfig(),
		esConfig:      newESConfig(),
	}
}

func (conf Config) GetPortNumber() string {
	return conf.port
}
