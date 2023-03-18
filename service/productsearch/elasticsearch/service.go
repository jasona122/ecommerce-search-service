package elasticsearch

import (
	"context"
	"fmt"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/contracts"

	"github.com/olivere/elastic/v7"
)

type Service interface {
	Search(ctx context.Context, query string, serviceAreaID string) ([]contracts.ESProductSearchResult, error)
}

type service struct {
	client *elastic.Client
}

func NewProductSearchESService(esConfig config.ElasticSearchConfig) (Service, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(esConfig.Host),
		elastic.SetHealthcheck(true),
	)

	if err != nil {
		return nil, fmt.Errorf("could not initialize elasticsearch client: #{err}", err)
	}

	return &service{
		client: client,
	}, nil
}

func (s service) Search(ctx context.Context, query string, serviceAreaID string) ([]contracts.ESProductSearchResult, error) {

}
