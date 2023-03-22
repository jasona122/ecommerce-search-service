package main

import (
	"fmt"

	"github.com/jasona122/ecommerce-search-service/config"
	"github.com/jasona122/ecommerce-search-service/elasticsearch"
	s "github.com/jasona122/ecommerce-search-service/server"
	"github.com/jasona122/ecommerce-search-service/service"
	"github.com/jasona122/ecommerce-search-service/service/productsearch"

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
	esConfigs := configs.GetElasticSearchConfig()
	productESService, err := elasticsearch.NewProductSearchESService(esConfigs)
	if err != nil {
		panic(fmt.Sprintf("could not initialize product ES service: %s", err))
	}
	productSearchService := productsearch.NewService(productESService)

	return service.Services{
		ProductSearchService: productSearchService,
	}
}
