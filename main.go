package main

import (
	"fmt"

	"github.com/jasona122/ecommerce-search-service/config"
	s "github.com/jasona122/ecommerce-search-service/server"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/codegangsta/negroni"
)

func main() {
	configs := config.Load()
	hystrix.Configure(config.GetHystrixLibraryConfig())

	server := negroni.New(negroni.NewRecovery())
	address := fmt.Sprintf(":%s", configs.GetPortNumber())

	router := s.NewRouter()
	server.UseHandler(router)

	server.Run(address)
}
