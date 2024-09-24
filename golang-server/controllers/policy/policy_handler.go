package policy

import (
	"fmt"
	"gPRC/dto"
	"gPRC/services/policy"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	policyService policy.Service
}

type Handler interface {
	GetPolicies(ctx *gin.Context)
}

func NewPolicyHandler(policyService policy.Service) Handler {
	return &handler{policyService: policyService}
}

func (h *handler) GetPolicies(ctx *gin.Context) {
	var queryRequest *dto.QueryRequest
	// Bind JSON request body to the PolicyRequest struct
	if err := ctx.ShouldBindJSON(&queryRequest); err != nil {
		fmt.Println("req {}", queryRequest)
		if err.Error() == "EOF" {
			queryRequest = nil
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	policies := h.policyService.GetPolicies(queryRequest)
	response := dto.QueryResponse{
		Total: len(policies),
		Data:  policies,
	}

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// }
	ctx.JSON(http.StatusOK, response)
}
