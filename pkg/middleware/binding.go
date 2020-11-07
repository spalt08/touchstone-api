package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// Bind is a middleware for reading and binding JSON body
func Bind(ctx *gin.Context, obj interface{}) *GenericAPIError {
	var body, readErr = ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()

	if readErr != nil {
		return NewBindingError(readErr)
	}

	// Make request body readable again
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var bindErr = jsoniter.Unmarshal(body, obj)
	if bindErr != nil {
		return NewBindingError(bindErr)
	}

	return nil
}
