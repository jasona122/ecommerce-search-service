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
	err := db.Init(configs.GetDatabaseConfig())
	if err != nil {
		panic(fmt.Sprintf("could not initialize DB: %s", err))
	}

	esService, err := elasticsearch.NewProductSearchESService(configs.GetElasticSearchConfig())
	if err != nil {
		panic(fmt.Sprintf("could not initialize product ES service: %s", err))
	}

	trendingService := trending.NewService(config.DatabaseConfig{})
	productSearchService := productsearch.NewService(esService, trendingService)
	shopSearchService := shopsearch.NewService(esService, trendingService)

	return service.Services{
		ProductSearchService: productSearchService,
		ShopSearchService:    shopSearchService,
		TrendingService:      trendingService,
	}
}
