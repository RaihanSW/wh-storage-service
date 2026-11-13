package response

import "wh-storage-service/app/model/entity"

type StorageResponse struct {
	GeneralResponse
}

type StorageDataResponse struct {
	GeneralResponse
	Data entity.Storage `json:"data,omitempty"`
}

type StorageListResponse struct {
	GeneralResponse
	Data []entity.Storage `json:"data,omitempty"`
}
