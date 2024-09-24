package register

import (
	"gPRC/services/register"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	registerService register.Service
}

func NewRegisterHandler(registerService register.Service) *handler {
	return &handler{registerService: registerService}
}

func (h *handler) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.registerService.Register())
}
