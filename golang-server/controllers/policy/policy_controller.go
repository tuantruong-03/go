package policy

import (
	"gPRC/bootstrap"

	"github.com/gin-gonic/gin"
)

type controller struct {
	ctx *bootstrap.Context
}

type Controller interface {
	RegisterRoutes(route *gin.RouterGroup)
}

func NewController(ctx *bootstrap.Context) Controller {
	return &controller{ctx: ctx}
}

func (controller *controller) RegisterRoutes(route *gin.RouterGroup) {
	policyHandler := NewPolicyHandler(controller.ctx.PolicyService)
	route.POST("", policyHandler.GetPolicies)
}
