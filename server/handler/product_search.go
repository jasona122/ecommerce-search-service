package handler

import (
	"fmt"
	"net/http"

	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/service/productsearch"
)

func ProductSearch(service productsearch.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		results, err := service.GetAllProducts(r.Context(), req)
		if err != nil {
			errorMessage := fmt.Sprintf("error in getting all products: %s", err)
			writeResponseJSON(w, contracts.InternalErrorResponse(errorMessage))
		}

		writeResponseJSON(w, contracts.Response{
			StatusCode: http.StatusOK,
			Data:       contracts.ProductSearchResponse{Results: results},
			Success:    true,
			Errors:     nil,
		})
	}
}

func transformRequest(r *http.Request) contracts.Request {
	return contracts.Request{
		Query:         getQueryFromQueryParam(r),
		ServiceAreaID: getServiceAreaIDFromHeader(r),
	}
}
