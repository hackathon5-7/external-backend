package repository

import (
	"app/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type StorageBillboard interface {
	AddBillboard(point models.Billboard) (int64, error)
	GetBillboardById(id int64) (models.Billboard, error)
	DeleteBillboardById(id int64) error
	GetAllBillboards() ([]models.Billboard, error)
	GetSizeStorageBillboards() (int, error)
	GetBillboardsBySectorId(sectorId int, limit int) ([]models.Billboard, error)
}

type StorageSectors interface {
	AddSector(point models.Sector) (int64, error)
	GetSectorForCoordinate(coordinate models.Billboard) (int, error)
	GetSizeStorageSectors() (int, error)
}

type StorageRequestsToTheModel interface {
	AddRequest(request models.Request) (int64, error)
	DeleteRequest(RequestId int64) error
}

type Repository struct {
	StorageBillboard
	StorageRequestsToTheModel
	StorageSectors
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		StorageBillboard:          NewStorageBillboardPostgres(db),
		StorageRequestsToTheModel: NewStorageRequestsToTheModelPostgres(db),
		StorageSectors:            NewStorageSectorsPostgres(db),
	}
}
