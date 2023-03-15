package main

import (
	"fmt"

	s "github.com/jasona122/ecommerce-search-service/server"

	"github.com/codegangsta/negroni"
)

func main() {
	server := negroni.New(negroni.NewRecovery())
	port := "8080"
	address := fmt.Sprintf(":%s", port)

	router := s.NewRouter()
	server.UseHandler(router)

	server.Run(address)
}
