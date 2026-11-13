package request

type StorageRequest struct {
	Name        string  `json:"name" validate:"required"`
	GoodsUID    *string `json:"goods_uid"`
	WarehouseID int     `json:"warehouse_id" validate:"required"`
	BlockID     int     `json:"block_id" validate:"required"`
	Status      int     `json:"status" validate:"required"`
	XRow        *int    `json:"x_row"`
	YRow        *int    `json:"y_row"`
}
