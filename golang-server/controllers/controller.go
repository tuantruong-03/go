package controllers

import (
	"gPRC/bootstrap"
	policyController "gPRC/controllers/policy"
	registerController "gPRC/controllers/register"

	"github.com/gin-gonic/gin"
)

func RegisterController(router *gin.Engine, ctx *bootstrap.Context) {
	policesGroup := router.Group("/api/policies")
	registerGroup := router.Group("/api/register")
	registerController.NewController(ctx).RegiserRoutes(registerGroup)
	policyController.NewController(ctx).RegisterRoutes(policesGroup)
}
