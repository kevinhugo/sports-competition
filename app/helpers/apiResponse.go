package helpers

import "net/http"

// "fmt"

type Response struct {
	Status      uint   `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Data        any    `json:"data,omitempty"`
	Meta        any    `json:"meta,omitempty"`
}

/*
General Response template.

# First data in additionalData will be considered as data objecct

Second data in additionalData will be consiedered as meta object
*/
func CreateResponse(status uint, message string, description string, additionalData ...any) Response {
	var responseData any
	if len(additionalData) > 0 {
		responseData = additionalData[0]
	}
	var responseMeta any
	if len(additionalData) > 1 {
		responseMeta = additionalData[1]
	}

	return Response{
		status,
		message,
		description,
		responseData,
		responseMeta,
	}
}

/*
General error response data, use this if there's an undefined error or as general error response
*/
func CreateSuccessResponse(additionalData ...any) Response {
	var additionalDataLength int = len(additionalData)
	var description string
	if additionalDataLength > 0 {
		description = additionalData[0].(string)
	} else {
		description = "Operation has been successfully executed."
	}
	var responseData any
	if len(additionalData) > 1 {
		responseData = additionalData[1]
	} else {
		responseData = map[string]any{}
	}
	var responseMeta any
	if len(additionalData) > 2 {
		responseMeta = additionalData[2]
	} else {
		responseMeta = map[string]any{}
	}
	return CreateResponse(http.StatusOK, "Success", description, responseData, responseMeta)
}

/*
General error response data, use this if there's an undefined error or as general error response
*/
func CreateGeneralErrorResponse(additionalData ...any) Response {
	var additionalDataLength int = len(additionalData)
	var description string
	if additionalDataLength > 0 {
		description = additionalData[0].(string)
	} else {
		description = "An error occured on server side, please try again later."
	}
	return CreateResponse(http.StatusInternalServerError, "Something went wrong", description)
}

/*
Bad Request response data, use this if response is unprocessable entity
*/
func CreateBadRequestResponse(additionalData ...any) Response {
	var additionalDataLength int = len(additionalData)
	var description string
	if additionalDataLength > 0 {
		description = additionalData[0].(string)
	} else {
		description = "Data cannot be processed."
	}
	return CreateResponse(http.StatusBadRequest, "Bad Request", description)
}

/*
Unauthorized response data, use this if response is unprocessable entity
*/
func CreateUnauthorizedResponse(additionalData ...any) Response {
	var additionalDataLength int = len(additionalData)
	var description string
	if additionalDataLength > 0 {
		description = additionalData[0].(string)
	} else {
		description = "Not permitted."
	}
	return CreateResponse(http.StatusUnauthorized, "Unauthorized", description)
}
