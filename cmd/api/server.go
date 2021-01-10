// Package api JSBench Server
//
// Backend REST server for jsbench
//
//     Schemes: http, https
//     Host: localhost:8081
//     BasePath: /v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//       jwtToken:
//         type: http
//         description: JWT token can be obtained by /login endpoint
//         scheme: Bearer
//         bearerFormat: JWT
//
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"

	"touchstone-api/external/postgres"
	"touchstone-api/pkg/middleware"
	"touchstone-api/pkg/user"
	"touchstone-api/pkg/utils/env"
)

func main() {
	var port = env.GetServingPort()
	var handler = gin.New()

	// Connect database
	var db = postgres.NewConnection(true)
	defer db.Close()

	// Setup services
	var userService = user.NewService(db)

	// Apply middlewares
	middleware.Setup(handler)

	// Apply routes
	user.Setup(handler, userService)

	// Run server
	handler.Run(port)
}
