package contracts

import "fmt"

type TrendingSchemaDB struct {
	ID            int    `json:"id" db:"id"`
	Query         string `json:"query" db:"query"`
	QueryCount    int    `json:"query_count" db:"query_count"`
	ServiceAreaID string `json:"service_area_id" db:"service_area_id"`
}

type GetTrendingServiceResult struct {
	Query      string `json:"query"`
	QueryCount int    `json:"count"`
}

type GetTrendingServiceResponse struct {
	Results []GetTrendingServiceResult `json:"results"`
}

func (response GetTrendingServiceResponse) DataMarker() {

}

type EditTrendingServiceResponse struct {
	Message string `json:"message"`
}

func (response EditTrendingServiceResponse) DataMarker() {

}

func DefaultEditSuccessResponse(query string, serviceAreaID string) EditTrendingServiceResponse {
	return EditTrendingServiceResponse{
		Message: fmt.Sprintf("Successful in editing query %s for Service Area ID %s", query, serviceAreaID),
	}
}
