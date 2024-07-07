package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
)

type RecomendationOutput struct {
	SectorId int                `json:"sector_id"`
	Value    int                `json:"value" db:"value"`
	Points   []models.Billboard `json:"points"`
}

type MLRequestService struct {
	reposSectors   repository.StorageSectors
	reposBillboard repository.StorageBillboard
}

func NewMLRequestService(reposSectors repository.StorageSectors, reposBillboard repository.StorageBillboard) *MLRequestService {
	return &MLRequestService{
		reposSectors:   reposSectors,
		reposBillboard: reposBillboard,
	}
}

func (s MLRequestService) GetRecomendation() ([]RecomendationOutput, error) {
	var data []RecomendationOutput

	countSectors, err := s.reposSectors.GetSizeStorageSectors()
	if err != nil {
		return data, err
	}

	for i := 1; i <= countSectors; i++ {
		billboard, err := s.reposBillboard.GetBillboardsBySectorId(i, 3)
		if err != nil {
			return nil, err
		}

		if billboard == nil {
			continue
		}

		data = append(data, RecomendationOutput{
			SectorId: i,
			Points:   billboard,
		})
	}

	return data, nil
}
