package repository

import (
	"app/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type StorageBillboard interface {
	AddBillboard(point models.Billboard) (int64, error)
	GetBillboardById(id int64) (models.Billboard, error)
	DeleteBillboardById(id int64) error
}

type StorageRequestsToTheModel interface {
	AddRequest(request models.Request) (int64, error)
	GetRequest(RequestId int64) (models.Request, error)
	DeleteRequest(RequestId int64) (int64, error)
}

type Repository struct {
	StorageBillboard
	StorageRequestsToTheModel
}

func NewRepository(db *sqlx.DB, rc *RedisClient) *Repository {
	return &Repository{
		StorageBillboard:          NewStorageBillboardPostgres(db),
		StorageRequestsToTheModel: NewStorageRequestsToTheModelRedis(rc),
	}
}
