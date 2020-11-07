package middleware

import (
	"github.com/gin-gonic/gin"
)

// GenericAPIResponse is a wrapper for all API responses
// swagger:model
type GenericAPIResponse struct {
	// Response notation version
	// example: 1
	Version int `json:"version"`

	// Request result
	// example: null
	Result interface{} `json:"result"`

	// Request error
	// example: null
	Error interface{} `json:"error"`
}

// ErrorAPIResponse is a wrapper for all API error responses
// swagger:model
type ErrorAPIResponse struct {
	GenericAPIResponse

	// Error details
	Error *GenericAPIError `json:"error"`
}

// RespondError sends error response with >200 status code with detailed error
func RespondError(ctx *gin.Context, err *GenericAPIError) {
	ctx.AbortWithStatusJSON(err.Code, &ErrorAPIResponse{
		GenericAPIResponse: GenericAPIResponse{
			Version: 1,
			Result:  nil,
		},
		Error: err,
	})
}

// RespondSuccess sends response with 200 status code with result body
func RespondSuccess(ctx *gin.Context, result interface{}) {
	ctx.JSON(200, &GenericAPIResponse{
		Version: 1,
		Result:  result,
		Error:   nil,
	})
}
