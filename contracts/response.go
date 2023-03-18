package contracts

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int            `json:"-"`
	Data       ResponseData   `json:"data"`
	Success    bool           `json:"success"`
	Errors     ResponseErrors `json:"errors"`
}

type ResponseData interface {
	DataMarker()
}

type ResponseErrors []ResponseError

type ResponseError struct {
	Message string `json:"message"`
}

type EmptyResponseData struct {
}

func (erd EmptyResponseData) DataMarker() {

}

func InternalErrorResponse(message string) Response {
	errorResponse := ResponseError{
		Message: message,
	}

	return Response{
		StatusCode: http.StatusInternalServerError,
		Data:       EmptyResponseData{},
		Success:    false,
		Errors:     ResponseErrors{errorResponse},
	}
}

func (response Response) ToJSON() ([]byte, error) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return []byte{}, err
	}

	return responseJSON, nil
}
