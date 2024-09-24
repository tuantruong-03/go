package register

import (
	"gPRC/bootstrap"

	"github.com/gin-gonic/gin"
)

type controller struct {
	ctx *bootstrap.Context
}

type Controller interface {
	RegiserRoutes(route *gin.RouterGroup)
}

func NewController(ctx *bootstrap.Context) Controller {
	return &controller{ctx: ctx}
}

func (c *controller) RegiserRoutes(route *gin.RouterGroup) {
	registerHandler := NewRegisterHandler(c.ctx.RegisterService)
	route.GET("/health", registerHandler.Register)
}
