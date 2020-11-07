package user_test

import (
	"testing"

	"jsbnch/external/postgres"
	"jsbnch/pkg/middleware"
	"jsbnch/pkg/user"
	"jsbnch/pkg/utils/tests"

	"github.com/stretchr/testify/assert"
)

func setupServer(t *testing.T) *tests.TestServer {
	var db = postgres.NewConnection()
	var service = user.NewService(db)
	var server = tests.NewTestServer(t)

	user.Setup(server.Handler, service)

	return server
}

func TestLoginValidation(t *testing.T) {
	endpoint := "/v1/login"
	body := "invalid body"

	server := setupServer(t)
	response := server.POST(endpoint, body)
	payload := &middleware.ErrorAPIResponse{}

	server.BindResponse(response, payload)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, payload.Error.Code, 400)
	assert.Equal(t, payload.Error.Message, "Bad Request")
}

func TestLoginSuccessful(t *testing.T) {
	endpoint := "/v1/login"
	body := `{"accessToken": "044d8fa5852a87528ed95e156c0c5010d162818c"}`

	server := setupServer(t)
	response := server.POST(endpoint, body)
	payload := &struct {
		Result user.LoginResponse
	}{}

	server.BindResponse(response, payload)

	assert.Equal(t, 200, response.Code)
	assert.Greater(t, len(payload.Result.Token), 0)
	assert.EqualValues(t, payload.Result.User.ID, 5869473)
	assert.Equal(t, payload.Result.User.Username, "spalt08")
	assert.EqualValues(t, payload.Result.User.Name, "Konstantin Darutkin")
}

func TestMeEndpoint(t *testing.T) {
	endpoint := "/v1/login"
	body := `{"accessToken": "044d8fa5852a87528ed95e156c0c5010d162818c"}`

	server := setupServer(t)
	response := server.POST(endpoint, body)
	payload := &struct {
		Result user.LoginResponse
	}{}

	server.BindResponse(response, payload)

	assert.Equal(t, 200, response.Code)
	assert.Greater(t, len(payload.Result.Token), 0)
	assert.EqualValues(t, payload.Result.User.ID, 5869473)
	assert.Equal(t, payload.Result.User.Username, "spalt08")
	assert.EqualValues(t, payload.Result.User.Name, "Konstantin Darutkin")
}
