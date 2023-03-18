package productsearch

import (
	"context"

	"github.com/jasona122/ecommerce-search-service/contracts"
)

type Service interface {
	GetAllProducts(ctx context.Context, req contracts.Request) ([]contracts.ProductSearchResult, error)
}

type service struct {
}

func NewService()
