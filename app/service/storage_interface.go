package service

import (
	"context"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
)

type StorageServiceInterface interface {
	StorageCreate(ctx context.Context, req request.StorageRequest) (entity.Storage, int, error)
	StorageGetAll(ctx context.Context) ([]entity.Storage, int, error)
	StorageGetByID(ctx context.Context, uid string) (entity.Storage, int, error)
	StorageUpdate(ctx context.Context, uid string, req request.StorageRequest) (int, error)
	StorageDelete(ctx context.Context, uid string) (int, error)
}
