// Package user contains all related logic, including authorization
package user

import (
	"github.com/gin-gonic/gin"
)

// Setup will attach controllers for related routes to gin instance
func Setup(router *gin.Engine, service *Service) {
	var routes = router.Group("/v1")
	var ctrl = &controller{
		service: service,
	}

	routes.GET("/me", ctrl.Me)
	routes.POST("/login/github", ctrl.GithubLogin)
}
