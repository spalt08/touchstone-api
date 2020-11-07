package user

import (
	"context"

	"jsbnch/pkg/middleware"
	"jsbnch/pkg/model"

	"github.com/go-pg/pg/v10"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Service is a wrapper for all user-related business login
type Service struct {
	db *pg.DB
}

// NewService is constructor function
func NewService(db *pg.DB) *Service {
	return &Service{
		db: db,
	}
}

// GetGithubInfo makes and API call to github and return user data by API call
func (svc *Service) GetGithubInfo(accessToken string) (*github.User, *middleware.GenericAPIError) {
	var ctx = context.Background()
	var tokenSource = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	var client = github.NewClient(oauth2.NewClient(ctx, tokenSource))
	var user, _, err = client.Users.Get(ctx, "")

	if err != nil {
		return nil, middleware.NewForbiddenError(err, "Forbidden")
	}

	return user, nil
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
