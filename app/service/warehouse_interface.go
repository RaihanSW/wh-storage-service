package service

import (
	"context"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
)

type WarehouseServiceInterface interface {
	WarehouseCreate(ctx context.Context, req request.WarehouseRequest) (entity.Warehouse, int, error)
	WarehouseGetAll(ctx context.Context) ([]entity.Warehouse, int, error)
	WarehouseGetByID(ctx context.Context, id int) (entity.Warehouse, int, error)
	WarehouseUpdate(ctx context.Context, id int, req request.WarehouseRequest) (int, error)
	WarehouseDelete(ctx context.Context, id int) (int, error)
}
