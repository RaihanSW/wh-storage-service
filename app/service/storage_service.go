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

type storageServiceImpl struct {
	cfg         *config.ApplicationConfig
	storageRepo repository.StorageRepositoryInterface
}

func NewStorageServiceImpl(
	cfg *config.ApplicationConfig,
	storageRepo repository.StorageRepositoryInterface) StorageServiceInterface {

	return &storageServiceImpl{
		cfg:         cfg,
		storageRepo: storageRepo,
	}
}

func (m *storageServiceImpl) StorageCreate(ctx context.Context, req request.StorageRequest) (entity.Storage, int, error) {
	data := entity.Storage{
		Name:        req.Name,
		GoodsUID:    req.GoodsUID,
		WarehouseID: req.WarehouseID,
		BlockID:     req.BlockID,
		Status:      req.Status,
		XRow:        req.XRow,
		YRow:        req.YRow,
	}

	uid, err := m.storageRepo.InsertStorage(ctx, data)
	if err != nil {
		return entity.Storage{}, http.StatusInternalServerError, fmt.Errorf("uc.StorageCreate (InsertStorage) : %v", err)
	}

	data.UID = uid
	return data, http.StatusCreated, nil
}

func (m *storageServiceImpl) StorageGetAll(ctx context.Context) ([]entity.Storage, int, error) {
	data, err := m.storageRepo.GetAllStorage(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.StorageGetAll (GetAllStorage) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *storageServiceImpl) StorageGetByID(ctx context.Context, uid string) (entity.Storage, int, error) {
	data, err := m.storageRepo.GetStorageByUID(ctx, uid)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.Storage{}, http.StatusNotFound, fmt.Errorf("uc.StorageGetByID : storage not found")
	case err != nil:
		return entity.Storage{}, http.StatusInternalServerError, fmt.Errorf("uc.StorageGetByID (GetStorageByUID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *storageServiceImpl) StorageUpdate(ctx context.Context, uid string, req request.StorageRequest) (int, error) {
	resultRows, err := m.storageRepo.UpdateStorageByUID(ctx, entity.Storage{
		UID:         uid,
		Name:        req.Name,
		GoodsUID:    req.GoodsUID,
		WarehouseID: req.WarehouseID,
		BlockID:     req.BlockID,
		Status:      req.Status,
		XRow:        req.XRow,
		YRow:        req.YRow,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.StorageUpdate (UpdateStorageByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.StorageUpdate : storage not found")
	}

	return http.StatusOK, nil
}

func (m *storageServiceImpl) StorageDelete(ctx context.Context, uid string) (int, error) {
	resultRows, err := m.storageRepo.DeleteStorageByUID(ctx, uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.StorageDelete (DeleteStorageByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.StorageDelete : storage not found")
	}

	return http.StatusOK, nil
}
