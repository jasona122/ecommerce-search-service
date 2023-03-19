package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/contracts"

	"github.com/olivere/elastic/v7"
)

type Service interface {
	Search(ctx context.Context, query string, serviceAreaID string) ([]contracts.ESProductSearchResultSource, error)
}

type service struct {
	client *elastic.Client
	config config.ElasticSearchConfig
}

func NewProductSearchESService(esConfig config.ElasticSearchConfig) (Service, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(esConfig.Host),
		elastic.SetHealthcheck(true),
	)

	if err != nil {
		return nil, fmt.Errorf("could not initialize elasticsearch client: %s", err)
	}

	return &service{
		client: client,
		config: esConfig,
	}, nil
}

func (s service) Search(ctx context.Context, query string, serviceAreaID string) ([]contracts.ESProductSearchResultSource, error) {
	esQuery := getQuery(ctx, query, serviceAreaID)
	searchResult, err := s.client.Search().
		Index(s.config.IndexName).
		Query(esQuery).
		Size(s.config.Size).
		Do(ctx)

	if err != nil {
		return []contracts.ESProductSearchResultSource{}, fmt.Errorf("could not get results from elasticsearch: %s", err)
	}

	var productResults []contracts.ESProductSearchResultSource
	var product contracts.ESProductSearchResultSource

	for _, hit := range searchResult.Hits.Hits {
		json.Unmarshal(hit.Source, &product)
		productResults = append(productResults, product)
	}

	return productResults, nil
}

func getQuery(ctx context.Context, query string, serviceAreaID string) elastic.Query {
	query = strings.ToLower(query)

	return elastic.
		NewBoolQuery().
		Filter(elastic.NewTermQuery("service_area_id", serviceAreaID)).
		Must(
			elastic.NewBoolQuery().
				Should(
					elastic.NewMatchQuery("name", query).Operator("AND").Fuzziness("AUTO"),
					elastic.NewMatchQuery("description", query).Operator("AND").Fuzziness("AUTO"),
				),
		)
}
