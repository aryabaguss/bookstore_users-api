package errors

import "net/http"

type RestErr struct {
	Message string `json:"messsage"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadrequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}