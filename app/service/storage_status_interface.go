package service

import (
	"context"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
)

type StorageStatusServiceInterface interface {
	StorageStatusCreate(ctx context.Context, req request.StorageStatusRequest) (entity.StorageStatus, int, error)
	StorageStatusGetAll(ctx context.Context) ([]entity.StorageStatus, int, error)
	StorageStatusGetByID(ctx context.Context, id int) (entity.StorageStatus, int, error)
	StorageStatusUpdate(ctx context.Context, id int, req request.StorageStatusRequest) (int, error)
	StorageStatusDelete(ctx context.Context, id int) (int, error)
}
