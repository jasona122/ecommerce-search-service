package shopsearch

import (
	"context"

	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/elasticsearch"
)

type Service interface {
	GetAllProductsFromShop(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error)
}

type service struct {
	esService elasticsearch.Service
}

func NewService(esService elasticsearch.Service) Service {
	return &service{
		esService: esService,
	}
}

func (s service) GetAllProductsFromShop(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error) {
	esResults, err := s.esService.SearchShops(ctx, req.Query, req.ServiceAreaID)
	if err != nil {
		return []contracts.ProductSearchResult{}, err
	}

	var searchResults []contracts.ProductSearchResult

	for _, result := range esResults {
		if result.Quantity <= 0 {
			continue
		}

		searchResults = append(searchResults, result.TransformToProductSearchResult())
	}

	return searchResults, nil
}
