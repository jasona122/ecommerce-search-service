package productsearch

import (
	"context"

	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/elasticsearch"
)

type Service interface {
	GetAllProducts(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error)
}

type service struct {
	esService elasticsearch.Service
}

func NewService(esService elasticsearch.Service) Service {
	return &service{
		esService: esService,
	}
}

func (s service) GetAllProducts(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error) {
	esResults, err := s.esService.SearchProducts(ctx, req.Query, req.ServiceAreaID)
	if err != nil {
		return []contracts.ProductSearchResult{}, err
	}

	var searchResults []contracts.ProductSearchResult

	for _, result := range esResults {
		searchResult := contracts.ProductSearchResult{
			Name:        result.Name,
			Category:    result.Category,
			Description: result.Description,
			ImageURL:    result.ImageURL,
			Price:       result.Price,
			Quantity:    result.Quantity,
			ShopName:    result.ShopName,
		}
		searchResults = append(searchResults, searchResult)
	}

	return searchResults, nil
}
