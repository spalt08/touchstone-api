package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"touchstone-api/pkg/middleware"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// TestServer is abstraction for more exlicit tests
type TestServer struct {
	Handler   *gin.Engine
	AuthToken string
	t         *testing.T
}

// NewTestServer creates a test gin server instance with actual middlewares
func NewTestServer(t *testing.T) *TestServer {
	gin.SetMode("release")

	var handler = gin.New()

	middleware.Setup(handler)

	return &TestServer{
		Handler:   handler,
		AuthToken: "",
		t:         t,
	}
}

// Request will make a request to test server endpoint
func (server *TestServer) Request(method string, endpoint string, body string) *httptest.ResponseRecorder {
	var response = httptest.NewRecorder()
	var buffer = bytes.NewBuffer([]byte(body))

	var request, _ = http.NewRequest(method, endpoint, buffer)

	if len(server.AuthToken) > 0 {
		request.Header.Set("Authorization", "bearer "+server.AuthToken)
	}

	server.Handler.ServeHTTP(response, request)

	return response
}

// POST request to test server endpoint
func (server *TestServer) POST(endpoint string, body string) *httptest.ResponseRecorder {
	return server.Request("POST", endpoint, body)
}

// GET request to test server endpoint
func (server *TestServer) GET(endpoint string) *httptest.ResponseRecorder {
	return server.Request("GET", endpoint, "")
}

// BindResponse buffer to generic response object
func (server *TestServer) BindResponse(response *httptest.ResponseRecorder, payload interface{}) {
	var responseBody, readErr = ioutil.ReadAll(response.Body)

	if readErr != nil {
		server.t.Fatal("Unable to read response buffer")
	}

	var bindErr = jsoniter.Unmarshal(responseBody, payload)
	if bindErr != nil {
		server.t.Fatal("Unable to bind response buffer")
	}
}
