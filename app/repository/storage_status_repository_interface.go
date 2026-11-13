package repository

import (
	"context"
	"wh-storage-service/app/model/entity"
)

type StorageStatusRepositoryInterface interface {
	InsertStorageStatus(ctx context.Context, data entity.StorageStatus) (int, error)
	GetAllStorageStatus(ctx context.Context) ([]entity.StorageStatus, error)
	GetStorageStatusByID(ctx context.Context, id int) (entity.StorageStatus, error)
	UpdateStorageStatusByID(ctx context.Context, data entity.StorageStatus) (int64, error)
	DeleteStorageStatusByID(ctx context.Context, id int) (int64, error)
}
