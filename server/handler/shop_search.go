package handler

import (
	"fmt"
	"net/http"

	"github.com/jasona122/ecommerce-search-service/service/shopsearch"
)

func ShopSearch(service shopsearch.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		results, err := service.GetAllProductsFromShop(r.Context(), req)
		if err != nil {
			errorMessage := fmt.Sprintf("error in getting products from shop for query %s: %s", req.Query, err)
			writeFailureResponse(w, http.StatusInternalServerError, errorMessage)
		} else {
			writeSuccessProductResponse(w, results)
		}
	}
}
