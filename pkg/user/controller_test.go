package user_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"touchstone-api/external/postgres"
	"touchstone-api/pkg/middleware"
	"touchstone-api/pkg/model"
	"touchstone-api/pkg/user"
	"touchstone-api/pkg/utils/httpclient"
	"touchstone-api/pkg/utils/tests"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

func setupServer(t *testing.T) (*tests.TestServer, *pg.Tx) {
	var db = postgres.NewConnection(true)
	var transaction, _ = db.Begin()

	var service = user.NewService(transaction)
	var api = tests.NewTestServer(t)

	user.Setup(api.Handler, service)

	return api, transaction
}

func TestLoginValidation(t *testing.T) {
	endpoint := "/v1/user/login"
	body := "invalid body"

	api, tx := setupServer(t)
	defer tx.Rollback()

	response := api.POST(endpoint, body)
	payload := &middleware.ErrorAPIResponse{}

	api.BindResponse(response, payload)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, payload.Error.Code, 400)
	assert.Equal(t, payload.Error.Message, "Bad Request")
}

func TestGithubLoginSuccessful(t *testing.T) {
	endpoint := "/v1/user/login"
	body := `{"code": "abcdef","state":"123456"}`

	// network mocks
	httpclient.SetMockHandler(func(req *http.Request) (*http.Response, error) {
		accessToken := "test_token_abc"
		switch req.URL.String() {
		case "https://github.com/login/oauth/access_token":
			body, err := ioutil.ReadAll(req.Body)

			assert.Nil(t, err)
			assert.Equal(t, req.Method, "POST")
			assert.Equal(t, string(body), `{"client_id":"56beab967846d0088cbe","client_secret":"9bd4390cb013e266d323a7f56552766c54399da8","code":"abcdef","state":"123456"}`)

			return httpclient.MockStringResponse(`{"access_token":"` + accessToken + `"}`)

		case "https://api.github.com/user":
			assert.Equal(t, req.Method, "GET")
			assert.Equal(t, req.Header.Get("Authorization"), "token "+accessToken)

			return httpclient.MockStringResponse(`{"id":5869473,"login":"spalt08","name":"Konstantin Darutkin","avatar_url":"http://github.com/test"}`)
		default:
			t.Fatal("Unexpected network request")
		}

		return nil, nil
	})

	api, tx := setupServer(t)
	defer tx.Rollback()

	response := api.POST(endpoint, body)
	payload := &struct {
		Result user.GithubLoginResponse
	}{}

	api.BindResponse(response, payload)

	assert.Equal(t, 200, response.Code)
	assert.Greater(t, len(payload.Result.Token), 0)
	assert.EqualValues(t, payload.Result.User.ID, 5869473)
	assert.Equal(t, payload.Result.User.Username, "spalt08")
	assert.EqualValues(t, payload.Result.User.Name, "Konstantin Darutkin")
	assert.NotNil(t, payload.Result.User.AvatarURL)
	assert.EqualValues(t, *(payload.Result.User.AvatarURL), "http://github.com/test")
}

func TestMeSuccessful(t *testing.T) {
	endpoint := "/v1/user/me"

	api, tx := setupServer(t)
	defer tx.Rollback()

	user, authToken := tests.CreateUser(tx, make(map[string]interface{}))
	api.AuthToken = authToken

	response := api.GET(endpoint)
	payload := &struct {
		Result *model.User
	}{}

	api.BindResponse(response, payload)

	assert.Equal(t, 200, response.Code)
	assert.EqualValues(t, payload.Result.ID, user.ID)
	assert.Equal(t, payload.Result.Username, user.Username)
	assert.EqualValues(t, payload.Result.Name, user.Name)
	assert.EqualValues(t, payload.Result.Email, user.Email)
	assert.EqualValues(t, payload.Result.AvatarURL, user.AvatarURL)
}

func TestMeWithoutAuthToken(t *testing.T) {
	endpoint := "/v1/user/me"

	api, _ := setupServer(t)

	response := api.GET(endpoint)
	payload := &struct {
		Error middleware.GenericAPIError
	}{}

	api.BindResponse(response, payload)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, payload.Error.Code, 400)
	assert.Equal(t, payload.Error.Message, "Bad Request")
}

func TestMeWithInvalidAuthToken(t *testing.T) {
	endpoint := "/v1/user/me"

	api, _ := setupServer(t)
	api.AuthToken = "invalid"

	response := api.GET(endpoint)
	payload := &struct {
		Error middleware.GenericAPIError
	}{}

	api.BindResponse(response, payload)

	assert.Equal(t, 401, response.Code)
	assert.Equal(t, payload.Error.Code, 401)
	assert.Equal(t, payload.Error.Message, "Unauthorized")
}
