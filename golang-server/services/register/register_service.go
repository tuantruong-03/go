package register

import "gPRC/repository"

type service struct {
	policyRepo repository.PolicyRepository
}

type Service interface {
	Register() string
}

func NewRegisterService() *service {
	return &service{}
}

func (s *service) Register() string {
	return "register successfully!"
}
