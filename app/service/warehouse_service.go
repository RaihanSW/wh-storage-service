package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
	"wh-storage-service/app/repository"
	"wh-storage-service/config"
)

type warehouseServiceImpl struct {
	cfg           *config.ApplicationConfig
	warehouseRepo repository.WarehouseRepositoryInterface
}

func NewWarehouseServiceImpl(
	cfg *config.ApplicationConfig,
	warehouseRepo repository.WarehouseRepositoryInterface) WarehouseServiceInterface {

	return &warehouseServiceImpl{
		cfg:           cfg,
		warehouseRepo: warehouseRepo,
	}
}

func (m *warehouseServiceImpl) WarehouseCreate(ctx context.Context, req request.WarehouseRequest) (entity.Warehouse, int, error) {
	data := entity.Warehouse{
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		TtlArea: req.TtlArea,
	}

	id, err := m.warehouseRepo.InsertWarehouse(ctx, data)
	if err != nil {
		return entity.Warehouse{}, http.StatusInternalServerError, fmt.Errorf("uc.WarehouseCreate (InsertWarehouse) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *warehouseServiceImpl) WarehouseGetAll(ctx context.Context) ([]entity.Warehouse, int, error) {
	data, err := m.warehouseRepo.GetAllWarehouse(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.WarehouseGetAll (GetAllWarehouse) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *warehouseServiceImpl) WarehouseGetByID(ctx context.Context, id int) (entity.Warehouse, int, error) {
	data, err := m.warehouseRepo.GetWarehouseByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.Warehouse{}, http.StatusNotFound, fmt.Errorf("uc.WarehouseGetByID : warehouse not found")
	case err != nil:
		return entity.Warehouse{}, http.StatusInternalServerError, fmt.Errorf("uc.WarehouseGetByID (GetWarehouseByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *warehouseServiceImpl) WarehouseUpdate(ctx context.Context, id int, req request.WarehouseRequest) (int, error) {
	resultRows, err := m.warehouseRepo.UpdateWarehouseByID(ctx, entity.Warehouse{
		ID:      id,
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		TtlArea: req.TtlArea,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.WarehouseUpdate (UpdateWarehouseByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.WarehouseUpdate : warehouse not found")
	}

	return http.StatusOK, nil
}

func (m *warehouseServiceImpl) WarehouseDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.warehouseRepo.DeleteWarehouseByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.WarehouseDelete (DeleteWarehouseByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.WarehouseDelete : warehouse not found")
	}

	return http.StatusOK, nil
}
