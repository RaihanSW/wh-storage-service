package repository

import (
	"context"
	"wh-storage-service/app/model/entity"
)

type WarehouseRepositoryInterface interface {
	InsertWarehouse(ctx context.Context, data entity.Warehouse) (int, error)
	GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error)
	GetWarehouseByID(ctx context.Context, id int) (entity.Warehouse, error)
	UpdateWarehouseByID(ctx context.Context, data entity.Warehouse) (int64, error)
	DeleteWarehouseByID(ctx context.Context, id int) (int64, error)
}
