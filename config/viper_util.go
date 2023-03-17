package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func getStringValue(key string) string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key: #{key} is not yet set"))
	}
	return viper.GetString(key)
}

func getIntValue(key string) int {
	value, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key: #{key} is not an integer value"))
	}
	return value
}
