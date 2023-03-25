package service

import (
	"github.com/jasona122/ecommerce-search-service/service/productsearch"
	"github.com/jasona122/ecommerce-search-service/service/shopsearch"
	"github.com/jasona122/ecommerce-search-service/service/trending"
)

type Services struct {
	ProductSearchService productsearch.Service
	ShopSearchService    shopsearch.Service
	TrendingService      trending.Service
}
