package handler

import (
	"fmt"
	"net/http"

	"github.com/jasona122/ecommerce-search-service/contracts"
	"github.com/jasona122/ecommerce-search-service/service/trending"
)

func GetTopTrendingQueries(service trending.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		results, err := service.GetTopTrendingQueries(r.Context(), req.ServiceAreaID)
		if err != nil {
			errorMessage := fmt.Sprintf("error in getting top trending data for service area id %s: %s", req.ServiceAreaID, err)
			writeFailureResponse(w, http.StatusInternalServerError, errorMessage)
		} else {
			trendingData := contracts.GetTrendingServiceResponse{Results: results}
			writeSuccessResponse(w, trendingData)
		}
	}
}

func AddTrendingQuery(service trending.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		response, err := service.AddTrendingQuery(r.Context(), req.Query, req.ServiceAreaID)

		if err != nil {
			errorMessage := fmt.Sprintf("error in adding new trending data for query %s with service area id %s: %s", req.Query, req.ServiceAreaID, err)
			writeFailureResponse(w, http.StatusInternalServerError, errorMessage)
		} else {
			writeSuccessResponse(w, response)
		}
	}
}

func DeleteTrendingQuery(service trending.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := transformRequest(r)

		response, err := service.DeleteTrendingQuery(r.Context(), req.Query, req.ServiceAreaID)

		if err != nil {
			errorMessage := fmt.Sprintf("error in deleting trending data for query %s with service area id %s: %s", req.Query, req.ServiceAreaID, err)
			writeFailureResponse(w, http.StatusInternalServerError, errorMessage)
		} else {
			writeSuccessResponse(w, response)
		}
	}
}
