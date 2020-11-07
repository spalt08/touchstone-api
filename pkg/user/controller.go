package user

import (
	"jsbnch/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (ctrl *controller) Login(ctx *gin.Context) {
	// swagger:operation POST /login User userLogin
	// ---
	// summary: Authentification
	// description: Can be used both for signing in and registration
	// parameters:
	// - name: Request
	//   in: body
	//   required: true
	//   schema:
	//     $ref: "#/definitions/LoginRequest"
	// responses:
	//   "200":
	//     description: "Success"
	//     schema:
	//       $ref: "#/definitions/LoginResponse"

	requestData := &LoginRequest{}
	err := middleware.Bind(ctx, requestData)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	githubUser, err := ctrl.service.GetGithubInfo(requestData.AccessToken)

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

	middleware.RespondSuccess(ctx, LoginResponse{
		User:  user,
		Token: token,
	})
}

func (ctrl *controller) Me(ctx *gin.Context) {
	// swagger:operation Get /me User userSelf
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
