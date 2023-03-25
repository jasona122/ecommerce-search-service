package config

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Driver   string
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Timeout  time.Duration
}

func newDBConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver:   getStringValue("DB_DRIVER"),
		Name:     getStringValue("DB_NAME"),
		Host:     getStringValue("DB_HOST"),
		Port:     getIntValue("DB_PORT"),
		User:     getStringValue("DB_USER"),
		Password: getStringValue("DB_PASSWORD"),
		Timeout:  time.Duration(getIntValue("DB_QUERY_TIMEOUT_MS")) * time.Millisecond,
	}
}

func (config DatabaseConfig) GetConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Name)
}

func (conf Config) GetDatabaseConfig() DatabaseConfig {
	return conf.dbConfig
}
