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

type blockLayoutServiceImpl struct {
	cfg            *config.ApplicationConfig
	blockLayoutRepo repository.BlockLayoutRepositoryInterface
}

func NewBlockLayoutServiceImpl(
	cfg *config.ApplicationConfig,
	blockLayoutRepo repository.BlockLayoutRepositoryInterface) BlockLayoutServiceInterface {

	return &blockLayoutServiceImpl{
		cfg:            cfg,
		blockLayoutRepo: blockLayoutRepo,
	}
}

func (m *blockLayoutServiceImpl) BlockLayoutCreate(ctx context.Context, req request.BlockLayoutRequest) (entity.BlockLayout, int, error) {
	data := entity.BlockLayout{
		WarehouseID: req.WarehouseID,
		BlockName:   req.BlockName,
		Floor:       req.Floor,
		XSum:        req.XSum,
		YSum:        req.YSum,
	}

	id, err := m.blockLayoutRepo.InsertBlockLayout(ctx, data)
	if err != nil {
		return entity.BlockLayout{}, http.StatusInternalServerError, fmt.Errorf("uc.BlockLayoutCreate (InsertBlockLayout) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *blockLayoutServiceImpl) BlockLayoutGetAll(ctx context.Context) ([]entity.BlockLayout, int, error) {
	data, err := m.blockLayoutRepo.GetAllBlockLayout(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.BlockLayoutGetAll (GetAllBlockLayout) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *blockLayoutServiceImpl) BlockLayoutGetByID(ctx context.Context, id int) (entity.BlockLayout, int, error) {
	data, err := m.blockLayoutRepo.GetBlockLayoutByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.BlockLayout{}, http.StatusNotFound, fmt.Errorf("uc.BlockLayoutGetByID : block layout not found")
	case err != nil:
		return entity.BlockLayout{}, http.StatusInternalServerError, fmt.Errorf("uc.BlockLayoutGetByID (GetBlockLayoutByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *blockLayoutServiceImpl) BlockLayoutUpdate(ctx context.Context, id int, req request.BlockLayoutRequest) (int, error) {
	resultRows, err := m.blockLayoutRepo.UpdateBlockLayoutByID(ctx, entity.BlockLayout{
		ID:          id,
		WarehouseID: req.WarehouseID,
		BlockName:   req.BlockName,
		Floor:       req.Floor,
		XSum:        req.XSum,
		YSum:        req.YSum,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.BlockLayoutUpdate (UpdateBlockLayoutByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.BlockLayoutUpdate : block layout not found")
	}

	return http.StatusOK, nil
}

func (m *blockLayoutServiceImpl) BlockLayoutDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.blockLayoutRepo.DeleteBlockLayoutByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.BlockLayoutDelete (DeleteBlockLayoutByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.BlockLayoutDelete : block layout not found")
	}

	return http.StatusOK, nil
}
