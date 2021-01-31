package user

import (
	"touchstone-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (ctrl *controller) GithubLogin(ctx *gin.Context) {
	// swagger:operation POST /login/github User userLogin
	// ---
	// summary: Authentification
	// description: Can be used both for signing in and registration
	// parameters:
	// - name: Request
	//   in: body
	//   required: true
	//   schema:
	//     $ref: "#/definitions/GithubLoginRequest"
	// responses:
	//   "200":
	//     description: "Success"
	//     schema:
	//       $ref: "#/definitions/GithubLoginResponse"

	requestData := &GithubLoginRequest{}
	err := middleware.Bind(ctx, requestData)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	accessToken, err := ctrl.service.GetGithuAccessToken(requestData.Code, requestData.State)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	githubUser, err := ctrl.service.GetGithubInfo(accessToken)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	user, err := ctrl.service.GetOrCreateUser(githubUser)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	token, err := middleware.CreateToken(user)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	middleware.RespondSuccess(ctx, GithubLoginResponse{
		User:  user,
		Token: token,
	})
}

func (ctrl *controller) Me(ctx *gin.Context) {
	// swagger:operation Get /me User userMe
	// ---
	// summary: Get Account
	// description: Returns current authorized user
	// security:
	// - jwtToken: []
	// responses:
	//   "200":
	//     description: "Success"
	//     schema:
	//       $ref: "#/definitions/User"

	userID, err := middleware.ParseToken(ctx)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	user, err := ctrl.service.GetUserByID(userID)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	middleware.RespondSuccess(ctx, user)
}
