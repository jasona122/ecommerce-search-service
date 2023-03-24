package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	port     string
	dbConfig DatabaseConfig
	esConfig ElasticSearchConfig
}

func Load() Config {
	viper.AutomaticEnv()

	return Config{
		port:     getStringValue("APP_PORT"),
		dbConfig: newDBConfig(),
		esConfig: newESConfig(),
	}
}

func (conf Config) GetPortNumber() string {
	return conf.port
}
