package main

import (
	"fmt"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/db"
	"github.com/jasona122/ecommerce-search-service/elasticsearch"
	s "github.com/jasona122/ecommerce-search-service/server"
	"github.com/jasona122/ecommerce-search-service/service"
	"github.com/jasona122/ecommerce-search-service/service/productsearch"
	"github.com/jasona122/ecommerce-search-service/service/shopsearch"
	"github.com/jasona122/ecommerce-search-service/service/trending"

	"github.com/codegangsta/negroni"
)

func main() {
	configs := config.Load()
	services := initServices(configs)

	server := negroni.New(negroni.NewRecovery())
	address := fmt.Sprintf(":%s", configs.GetPortNumber())

	router := s.NewRouter(services)
	server.UseHandler(router)

	server.Run(address)
}

func initServices(configs config.Config) service.Services {
	db.Init(config.DatabaseConfig{})

	esService, err := elasticsearch.NewProductSearchESService(configs.GetElasticSearchConfig())
	if err != nil {
		panic(fmt.Sprintf("could not initialize product ES service: %s", err))
	}
	productSearchService := productsearch.NewService(esService)
	shopSearchService := shopsearch.NewService(esService)
	trendingService := trending.NewService(config.DatabaseConfig{})

	return service.Services{
		ProductSearchService: productSearchService,
		ShopSearchService:    shopSearchService,
		TrendingService:      trendingService,
	}
}
