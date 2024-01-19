package rest

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    code,
		Err:     err,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "internal_server_error",
		Causes:  nil,
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "not_found",
		Causes:  nil,
	}
}
