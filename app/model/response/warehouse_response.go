package response

import "wh-storage-service/app/model/entity"

type WarehouseResponse struct {
	GeneralResponse
}

type WarehouseDataResponse struct {
	GeneralResponse
	Data entity.Warehouse `json:"data,omitempty"`
}

type WarehouseListResponse struct {
	GeneralResponse
	Data []entity.Warehouse `json:"data,omitempty"`
}
