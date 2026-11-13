package response

import "wh-storage-service/app/model/entity"

type BlockLayoutResponse struct {
	GeneralResponse
}

type BlockLayoutDataResponse struct {
	GeneralResponse
	Data entity.BlockLayout `json:"data,omitempty"`
}

type BlockLayoutListResponse struct {
	GeneralResponse
	Data []entity.BlockLayout `json:"data,omitempty"`
}
