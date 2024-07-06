package repository

import (
	"app/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type StorageRequestsToTheModelPostgres struct {
	db *sqlx.DB
}

func NewStorageRequestsToTheModelPostgres(db *sqlx.DB) *StorageRequestsToTheModelPostgres {
	return &StorageRequestsToTheModelPostgres{
		db: db,
	}
}

func (r StorageRequestsToTheModelPostgres) AddRequest(request models.Request) (int64, error) {
	return 0, nil
}

func (r StorageRequestsToTheModelPostgres) GetRequestByUserId(UserId int64) ([]models.Request, error) {
	return nil, nil
}

func (r StorageRequestsToTheModelPostgres) DeleteRequest(RequestId int64) (int64, error) {
	return 0, nil
}
