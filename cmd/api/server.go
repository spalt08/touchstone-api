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

	"jsbnch/external/postgres"
	"jsbnch/pkg/middleware"
	"jsbnch/pkg/user"
	"jsbnch/pkg/utils/env"
)

func main() {
	var port = env.GetServingPort()
	var handler = gin.New()

	// Connect database
	var db = postgres.NewConnection()
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
