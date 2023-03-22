package service

import (
	"github.com/jasona122/ecommerce-search-service/service/productsearch"
	"github.com/jasona122/ecommerce-search-service/service/shopsearch"
)

type Services struct {
	ProductSearchService productsearch.Service
	ShopSearchService    shopsearch.Service
}
