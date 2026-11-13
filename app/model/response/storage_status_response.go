package response

import "wh-storage-service/app/model/entity"

type StorageStatusResponse struct {
	GeneralResponse
}

type StorageStatusDataResponse struct {
	GeneralResponse
	Data entity.StorageStatus `json:"data,omitempty"`
}

type StorageStatusListResponse struct {
	GeneralResponse
	Data []entity.StorageStatus `json:"data,omitempty"`
}
