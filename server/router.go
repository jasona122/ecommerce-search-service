package server

import (
	"net/http"

	"github.com/jasona122/ecommerce-search-service/server/handler"
	"github.com/jasona122/ecommerce-search-service/service"

	"github.com/gorilla/mux"
)

func NewRouter(services service.Services) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler)
	router.HandleFunc("/products", handler.ProductSearch(services.ProductSearchService)).Methods(http.MethodGet)

	return router
}
