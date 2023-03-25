package handler

import (
	"net/http"

	"github.com/jasona122/ecommerce-search-service/contracts"
)

func getQueryFromQueryParam(r *http.Request) string {
	return r.URL.Query().Get("query")
}

func getCategoryFromQueryParam(r *http.Request) string {
	return r.URL.Query().Get("category")
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

func writeSuccessResponse(w http.ResponseWriter, data contracts.ResponseData) {
	writeResponseJSON(w, contracts.Response{
		StatusCode: http.StatusOK,
		Data:       data,
		Success:    true,
		Errors:     nil,
	})
}

func writeSuccessProductResponse(w http.ResponseWriter, results []contracts.ProductSearchResult) {
	data := contracts.ProductSearchResponse{Results: results}
	writeSuccessResponse(w, data)
}

func writeFailureResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := contracts.ResponseError{
		Message: message,
	}

	writeResponseJSON(w, contracts.Response{
		StatusCode: statusCode,
		Data:       contracts.EmptyResponseData{},
		Success:    false,
		Errors:     contracts.ResponseErrors{errorResponse},
	})
}

func transformRequest(r *http.Request) contracts.Request {
	return contracts.Request{
		Query:         getQueryFromQueryParam(r),
		Category:      getCategoryFromQueryParam(r),
		ServiceAreaID: getServiceAreaIDFromHeader(r),
	}
}
