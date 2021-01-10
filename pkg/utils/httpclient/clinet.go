// Package httpclient used as abstraction for arbitrary network HTTP requests for tests
package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client can mocked in tests
	Client HTTPClient = &http.Client{}

	// MockedResponse can be defined inside test file or SetMockHandler can be used
	MockedResponse func(req *http.Request) (*http.Response, error)
)

// MockClient is used for tests
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is a wrapper for mocked responses
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return MockedResponse(req)
}

// SetMockHandler is used to define mocked response
func SetMockHandler(handler func(req *http.Request) (*http.Response, error)) {
	Client = &MockClient{}
	MockedResponse = handler
}

// MockStringResponse is a helper function for creating responses
func MockStringResponse(body string) (*http.Response, error) {
	bodyBuffer := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	return &http.Response{
		StatusCode: 200,
		Body:       bodyBuffer,
	}, nil
}
