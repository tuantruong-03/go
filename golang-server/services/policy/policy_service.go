package policy

import (
	"gPRC/dto"
	"gPRC/models"
	"gPRC/repository"
)

type service struct {
	policyRepo repository.PolicyRepository
}

type Service interface {
	Save(policyReq dto.PolicyRequest) models.Policy
	GetPolicies(queryRequest *dto.QueryRequest) []models.Policy
}

func NewService(policyRepo repository.PolicyRepository) Service {

	return &service{policyRepo: policyRepo}
}

func (s *service) Save(policyReq dto.PolicyRequest) models.Policy {
	return models.Policy{}
}

func (s *service) GetPolicies(queryRequest *dto.QueryRequest) []models.Policy {
	return s.policyRepo.GetPolicies(queryRequest)
}
