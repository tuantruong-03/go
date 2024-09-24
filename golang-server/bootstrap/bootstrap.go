package bootstrap

import (
	"gPRC/repository"
	"gPRC/services/policy"
	"gPRC/services/register"
)

type Context struct {
	PolicyService   policy.Service
	RegisterService register.Service
}

func BuildApplicationContext() *Context {
	appContext := Context{}

	policyRepository := repository.NewPolicyRepository()

	// Left: "Service" interface , Right: "service" struct implements "Service" interface
	appContext.PolicyService = policy.NewService(policyRepository)
	appContext.RegisterService = register.NewRegisterService()
	return &appContext
}
