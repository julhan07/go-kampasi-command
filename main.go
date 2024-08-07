package command

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}

func NewSuccessWithMetaResponse(status int, data interface{}, meta interface{}) *Response {
	return &Response{
		Status:  status,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
		Meta:    meta,
	}
}

func NewErrorResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func NewCreatedResponse(status int, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	}
}

func NewBadRequestResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func NewNotFoundResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func NewUnauthorizedResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func NewInternalServerErrorResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}
