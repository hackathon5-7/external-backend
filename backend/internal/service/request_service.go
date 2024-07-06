package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
)

type RequestInput struct {
	AgeFrom int         `json:"age_from"`
	AgeTo   int         `json:"age_to"`
	Gender  string      `json:"gender"`
	Income  IncomeInput `json:"income"`
}

type IncomeInput struct {
	A bool `json:"a"`
	B bool `json:"b"`
	C bool `json:"c"`
}

type RequestService struct {
	repos *repository.Repository
}

func NewRequestService(repos *repository.Repository) *RequestService {
	return &RequestService{
		repos: repos,
	}
}

func (r *RequestService) GetValueForRequest(request RequestInput) (string, error) {
	return "", nil
}

func (r *RequestService) ProcessingRequest(request RequestInput) (string, error) {
	return "", nil
}

func (r *RequestService) AddRequest(request models.Request) (int64, error) {
	return 0, nil
}

func (r *RequestService) DeleteRequest(RequestId int64) error {
	return nil
}
