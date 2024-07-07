package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
)

type PointService struct {
	repos repository.StorageBillboard
}

func NewPointsService(repos repository.StorageBillboard) *PointService {
	return &PointService{
		repos: repos,
	}
}

func (s PointService) GetAllPoints() ([]models.Billboard, error) {
	return s.repos.GetAllBillboards()
}
