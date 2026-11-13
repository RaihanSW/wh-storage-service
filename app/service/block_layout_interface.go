package service

import (
	"context"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
)

type BlockLayoutServiceInterface interface {
	BlockLayoutCreate(ctx context.Context, req request.BlockLayoutRequest) (entity.BlockLayout, int, error)
	BlockLayoutGetAll(ctx context.Context) ([]entity.BlockLayout, int, error)
	BlockLayoutGetByID(ctx context.Context, id int) (entity.BlockLayout, int, error)
	BlockLayoutUpdate(ctx context.Context, id int, req request.BlockLayoutRequest) (int, error)
	BlockLayoutDelete(ctx context.Context, id int) (int, error)
}
