package handler

import (
	"fmt"
	"net/http"

	"github.com/jasona122/ecommerce-search-service/service/productsearch"
)

func ProductSearch(service productsearch.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		results, err := service.GetProducts(r.Context(), req)
		if err != nil {
			errorMessage := fmt.Sprintf("error in getting products for query %s: %s", req.Query, err)
			writeFailureResponse(w, http.StatusInternalServerError, errorMessage)
		} else {
			writeSuccessProductResponse(w, results)
		}
	}
}
