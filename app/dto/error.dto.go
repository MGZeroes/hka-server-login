package dto

import (
	"encoding/json"
)

// ErrorResponse represents the JSON structure for an error response
type ErrorResponse struct {
	Error errorStruct `json:"error"`
}

type errorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JsonError(code int, err error) string {
	return NewErrorResponse(code, err).ToJsonString()
}

func NewErrorResponse(code int, err error) *ErrorResponse {
	errorResponse := &ErrorResponse{Error: errorStruct{
		Code:    code,
		Message: err.Error(),
	}}
	return errorResponse
}

func (errorResp *ErrorResponse) ToJsonString() string {
	response, _ := json.Marshal(errorResp)
	return string(response)
}
