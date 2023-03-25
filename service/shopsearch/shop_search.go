package shopsearch

import (
	"context"
	"fmt"

	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/elasticsearch"
	"github.com/jasona122/ecommerce-search-service/service/trending"
)

type Service interface {
	GetAllProductsFromShop(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error)
}

type service struct {
	esService       elasticsearch.Service
	trendingService trending.Service
}

func NewService(esService elasticsearch.Service, trendingService trending.Service) Service {
	return &service{
		esService:       esService,
		trendingService: trendingService,
	}
}

func (s service) GetAllProductsFromShop(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error) {
	esResults, err := s.esService.SearchShops(ctx, req.Query, req.ServiceAreaID)
	if err != nil {
		return []contracts.ProductSearchResult{}, err
	}

	_, err = s.trendingService.IncrementQueryCount(ctx, req.Query, req.ServiceAreaID)
	if err != nil {
		fmt.Printf("unable to increment trending query count for query: %s in service area ID %s: %s\n", req.Query, req.ServiceAreaID, err)
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
