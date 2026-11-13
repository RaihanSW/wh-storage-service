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

type storageStatusServiceImpl struct {
	cfg               *config.ApplicationConfig
	storageStatusRepo repository.StorageStatusRepositoryInterface
}

func NewStorageStatusServiceImpl(
	cfg *config.ApplicationConfig,
	storageStatusRepo repository.StorageStatusRepositoryInterface) StorageStatusServiceInterface {

	return &storageStatusServiceImpl{
		cfg:               cfg,
		storageStatusRepo: storageStatusRepo,
	}
}

func (m *storageStatusServiceImpl) StorageStatusCreate(ctx context.Context, req request.StorageStatusRequest) (entity.StorageStatus, int, error) {
	data := entity.StorageStatus{Name: req.Name}

	id, err := m.storageStatusRepo.InsertStorageStatus(ctx, data)
	if err != nil {
		return entity.StorageStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.StorageStatusCreate (InsertStorageStatus) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *storageStatusServiceImpl) StorageStatusGetAll(ctx context.Context) ([]entity.StorageStatus, int, error) {
	data, err := m.storageStatusRepo.GetAllStorageStatus(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.StorageStatusGetAll (GetAllStorageStatus) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *storageStatusServiceImpl) StorageStatusGetByID(ctx context.Context, id int) (entity.StorageStatus, int, error) {
	data, err := m.storageStatusRepo.GetStorageStatusByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.StorageStatus{}, http.StatusNotFound, fmt.Errorf("uc.StorageStatusGetByID : storage status not found")
	case err != nil:
		return entity.StorageStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.StorageStatusGetByID (GetStorageStatusByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *storageStatusServiceImpl) StorageStatusUpdate(ctx context.Context, id int, req request.StorageStatusRequest) (int, error) {
	resultRows, err := m.storageStatusRepo.UpdateStorageStatusByID(ctx, entity.StorageStatus{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.StorageStatusUpdate (UpdateStorageStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.StorageStatusUpdate : storage status not found")
	}

	return http.StatusOK, nil
}

func (m *storageStatusServiceImpl) StorageStatusDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.storageStatusRepo.DeleteStorageStatusByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.StorageStatusDelete (DeleteStorageStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.StorageStatusDelete : storage status not found")
	}

	return http.StatusOK, nil
}
