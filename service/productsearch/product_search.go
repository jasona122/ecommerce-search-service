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
	productCategory := ""
	if contracts.IsValidProductCategory(req.Category) {
		productCategory = req.Category
	}

	esResults, err := s.esService.SearchProducts(ctx, req.Query, productCategory, req.ServiceAreaID)
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
