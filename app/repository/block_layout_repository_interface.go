package repository

import (
	"context"
	"wh-storage-service/app/model/entity"
)

type BlockLayoutRepositoryInterface interface {
	InsertBlockLayout(ctx context.Context, data entity.BlockLayout) (int, error)
	GetAllBlockLayout(ctx context.Context) ([]entity.BlockLayout, error)
	GetBlockLayoutByID(ctx context.Context, id int) (entity.BlockLayout, error)
	UpdateBlockLayoutByID(ctx context.Context, data entity.BlockLayout) (int64, error)
	DeleteBlockLayoutByID(ctx context.Context, id int) (int64, error)
}
