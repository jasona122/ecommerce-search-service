package config

type ElasticSearchConfig struct {
	Host      string
	IndexName string
}

func newESConfig() ElasticSearchConfig {
	return ElasticSearchConfig{
		Host:      getStringValue("ES_HOST"),
		IndexName: getStringValue("ES_INDEX_NAME"),
	}
}

func (conf Config) GetElasticSearchConfig() ElasticSearchConfig {
	return conf.esConfig
}
