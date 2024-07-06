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

type DatasetParser interface {
	GetBillboards(pathFile string) error
}

type SectorParser interface {
	GetSectors(pathFile string) error
	GetSectorsArray(pathFile string) ([]models.Sector, error)
}

type MLRequest interface {
	GetRecomendation() ([]RecomendationOutput, error)
}

type Service struct {
	Request
	Points
	DatasetParser
	SectorParser
	MLRequest
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Request:       NewRequestService(repos),
		Points:        NewPointsService(repos.StorageBillboard),
		DatasetParser: NewDatasetParserJSONService(repos.StorageBillboard, repos.StorageSectors),
		SectorParser:  NewSectorParserJSONService(repos.StorageSectors),
		MLRequest:     NewMLRequestService(repos.StorageSectors, repos.StorageBillboard),
	}
}
