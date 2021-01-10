package user

import (
	"bytes"
	"encoding/json"
	"net/http"

	"touchstone-api/external/postgres"
	"touchstone-api/pkg/middleware"
	"touchstone-api/pkg/model"
	"touchstone-api/pkg/utils/env"
	"touchstone-api/pkg/utils/httpclient"

	"github.com/go-pg/pg/v10"
	"github.com/google/go-github/github"
)

// Service is a wrapper for all user-related business login
type Service struct {
	db postgres.Connection
}

// GithubAccessTokenRequest is a payload for POST https://github.com/login/oauth/access_token
// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps#2-users-are-redirected-back-to-your-site-by-github
type GithubAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	State        string `json:"state"`
}

// GithubAccessTokenResponse is a response from POST https://github.com/login/oauth/access_token
// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps#2-users-are-redirected-back-to-your-site-by-github
type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// NewService is constructor function
func NewService(db postgres.Connection) *Service {
	return &Service{
		db: db,
	}
}

// GetGithuAccessToken makes and API call to github and return the access token
// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps#2-users-are-redirected-back-to-your-site-by-github
func (svc *Service) GetGithuAccessToken(code string, state string) (string, *middleware.GenericAPIError) {
	payload := &GithubAccessTokenRequest{
		ClientID:     env.GetGithubClientID(),
		ClientSecret: env.GetGithubClientSecret(),
		Code:         code,
		State:        state,
	}

	requestBody, err := json.Marshal(payload)

	if err != nil {
		return "", middleware.NewInternalError(err)
	}

	request, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestBody))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "", middleware.NewInternalError(err)
	}

	response, err := httpclient.Client.Do(request)

	if err != nil {
		return "", middleware.NewInternalError(err)
	}

	defer response.Body.Close()

	responseBody := &GithubAccessTokenResponse{}
	err = json.NewDecoder(response.Body).Decode(responseBody)

	if err != nil {
		return "", middleware.NewInternalError(err)
	}

	return responseBody.AccessToken, nil
}

// GetGithubInfo makes and API call to github and return user data by API call
// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps#3-use-the-access-token-to-access-the-api
func (svc *Service) GetGithubInfo(accessToken string) (*github.User, *middleware.GenericAPIError) {
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "token "+accessToken)

	if err != nil {
		return nil, middleware.NewInternalError(err)
	}

	response, err := httpclient.Client.Do(request)

	if err != nil {
		return nil, middleware.NewInternalError(err)
	}

	defer response.Body.Close()

	responseBody := &github.User{}
	err = json.NewDecoder(response.Body).Decode(responseBody)

	if err != nil {
		return nil, middleware.NewInternalError(err)
	}

	return responseBody, nil
}

// GetOrCreateUser will get or create user inside DB
func (svc *Service) GetOrCreateUser(data *github.User) (*model.User, *middleware.GenericAPIError) {
	user := &model.User{ID: *data.ID}
	err := svc.db.
		Model(user).
		WherePK().
		Select(user)

	if err == pg.ErrNoRows {
		user = &model.User{
			ID:        *data.ID,
			Username:  *data.Login,
			Name:      *data.Name,
			Email:     data.Email,
			Company:   data.Company,
			AvatarURL: data.AvatarURL,
		}
		_, err = svc.db.
			Model(user).
			Insert()
	}

	if err != nil {
		return nil, middleware.NewDatabaseError(err)
	}

	return user, nil
}

// GetUserByID will get the user
func (svc *Service) GetUserByID(userID int64) (*model.User, *middleware.GenericAPIError) {
	user := &model.User{ID: userID}
	err := svc.db.
		Model(user).
		WherePK().
		Select(user)

	if err != nil {
		return nil, middleware.NewDatabaseError(err)
	}

	return user, nil
}
