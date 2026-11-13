package repository

import (
	"context"
	"wh-storage-service/app/model/entity"
)

type StorageRepositoryInterface interface {
	InsertStorage(ctx context.Context, data entity.Storage) (string, error)
	GetAllStorage(ctx context.Context) ([]entity.Storage, error)
	GetStorageByUID(ctx context.Context, uid string) (entity.Storage, error)
	UpdateStorageByUID(ctx context.Context, data entity.Storage) (int64, error)
	DeleteStorageByUID(ctx context.Context, uid string) (int64, error)
}
