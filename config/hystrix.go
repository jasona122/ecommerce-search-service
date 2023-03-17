package config

import (
	"github.com/afex/hystrix-go/hystrix"
)

const SearchCommand = "SearchCommand"

func GetHystrixLibraryConfig() map[string]hystrix.CommandConfig {
	return map[string]hystrix.CommandConfig{
		SearchCommand: newHystrixConfig(),
	}
}

func newHystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                getIntValue("HYSTRIX_TIMEOUT_MS"),
		MaxConcurrentRequests:  getIntValue("HYSTRIX_MAX_CONCURRENT_REQUESTS"),
		RequestVolumeThreshold: getIntValue("HYSTRIX_REQUEST_VOLUME_THRESHOLD"),
		SleepWindow:            getIntValue("HYSTRIX_SLEEP_WINDOW_MS"),
		ErrorPercentThreshold:  getIntValue("HYSTRIX_ERROR_THRESHOLD_PERCENTAGE"),
	}
}
