package handler

import (
	"net/http"

	"github.com/jasona122/ecommerce-search-service/contracts"
)

func getQueryFromQueryParam(r *http.Request) string {
	return r.URL.Query().Get("query")
}

func getServiceAreaIDFromHeader(r *http.Request) string {
	return r.Header.Get("Service-Area-ID")
}

func writeResponseJSON(w http.ResponseWriter, response contracts.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	responseJSON, _ := response.ToJSON()
	w.Write(responseJSON)
}

func transformRequest(r *http.Request) contracts.Request {
	return contracts.Request{
		Query:         getQueryFromQueryParam(r),
		ServiceAreaID: getServiceAreaIDFromHeader(r),
	}
}
