package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

// GenericAPIError is a wrapper for all API errors
// swagger:model
type GenericAPIError struct {

	// Error code, often same as HTTP status code
	Code int `json:"code"`

	// Humanized message
	Message string `json:"message"`

	// Error details
	// example: null
	Data interface{} `json:"data"`

	Error error `json:"-"`
}

// NewBindingError transforms gin validation error to GenericApiError
func NewBindingError(err error) *GenericAPIError {
	var fields, _ = err.(validator.ValidationErrors)
	var data = make([]interface{}, len(fields))

	for i, field := range fields {
		data[i] = gin.H{
			"key":     field.Field(),
			"tag":     field.Tag(),
			"type":    field.Type().String(),
			"message": convertErrorMessage(field.Translate(nil)),
		}
	}

	return &GenericAPIError{
		Code:    400,
		Message: "Bad Request",
		Data:    data,
		Error:   err,
	}
}

// NewForbiddenError returns any 403 error
func NewForbiddenError(err error, message string) *GenericAPIError {
	return &GenericAPIError{
		Code:    403,
		Message: message,
		Error:   err,
	}
}

// NewDatabaseError is a wrapper for all possible database errors
func NewDatabaseError(err error) *GenericAPIError {
	return &GenericAPIError{
		Code:    500,
		Message: "Internal Server Error",
		Error:   err,
	}
}

// NewUnauthorizedError is a wrapper for auth-related issues
func NewUnauthorizedError(err error) *GenericAPIError {
	return &GenericAPIError{
		Code:    401,
		Message: "Unauthorized",
		Error:   err,
	}
}

// NewInternalError is a wrapper for internal unexpected issues
func NewInternalError(err error) *GenericAPIError {
	return &GenericAPIError{
		Code:    500,
		Message: "Internal Server Error",
		Error:   err,
	}
}

func convertErrorMessage(message string) string {
	var splitBy = "Error:"
	var index = strings.LastIndex(message, "Error:") + len(splitBy)

	if index == len(splitBy)-1 {
		index = 0
	}

	return message[index:]
}
