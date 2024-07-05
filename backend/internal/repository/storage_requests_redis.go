package repository

import "app/backend/internal/models"

type StorageRequestsToTheModelRedis struct {
	rc *RedisClient
}

func NewStorageRequestsToTheModelRedis(rc *RedisClient) *StorageRequestsToTheModelRedis {
	return &StorageRequestsToTheModelRedis{rc: rc}
}

func (r StorageRequestsToTheModelRedis) AddRequest(request models.Request) (int64, error) {
	return 0, nil
}

func (r StorageRequestsToTheModelRedis) GetRequest(RequestId int64) (models.Request, error) {
	return models.Request{}, nil
}

func (r StorageRequestsToTheModelRedis) DeleteRequest(RequestId int64) (int64, error) {
	return 0, nil
}
