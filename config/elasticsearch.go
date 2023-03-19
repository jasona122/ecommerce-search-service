package config

type ElasticSearchConfig struct {
	Host      string
	IndexName string
	Size      int
}

func newESConfig() ElasticSearchConfig {
	return ElasticSearchConfig{
		Host:      getStringValue("ES_HOST"),
		IndexName: getStringValue("ES_INDEX_NAME"),
		Size:      getIntValue("ES_MAX_RETURN_SIZE"),
	}
}

func (conf Config) GetElasticSearchConfig() ElasticSearchConfig {
	return conf.esConfig
}
