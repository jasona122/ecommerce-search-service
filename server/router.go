package server

import (
	"github.com/jasona122/ecommerce-search-service/server/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler)

	return router
}
