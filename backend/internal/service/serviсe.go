package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
)

type Request interface {
	GetValueForRequest(request RequestInput) (string, error)
	ProcessingRequest(request RequestInput) (string, error)
	AddRequest(request models.Request) (int64, error)
	DeleteRequest(RequestId int64) error
}

type Points interface {
	GetAllPoints() ([]models.Billboard, error)
}

type Service struct {
	Request
	Points
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Request: NewRequestService(repos),
		Points:  NewPointsService(repos.StorageBillboard),
	}
}
