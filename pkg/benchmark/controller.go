package benchmark

import (
	"touchstone-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type controller struct {
}

func (ctrl *controller) Create(ctx *gin.Context) {
	// swagger:operation POST /benchmark Benchmark benchmarkCreate
	// ---
	// summary: Benchmark creation
	// parameters:
	// - name: Request
	//   in: body
	//   required: true
	//   schema:
	//     $ref: "#/definitions/CreateRequest"
	// responses:
	//   "200":
	//     description: "Success"
	//     schema:
	//       $ref: "#/definitions/CreateRequest"

	userID, err := middleware.ParseToken(ctx)

	if err != nil {
		middleware.RespondError(ctx, err)
		return
	}

	middleware.RespondSuccess(ctx, userID)
}
