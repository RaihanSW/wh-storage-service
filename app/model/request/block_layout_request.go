package request

type BlockLayoutRequest struct {
	WarehouseID int     `json:"warehouse_id" validate:"required"`
	BlockName   string  `json:"block_name" validate:"required"`
	Floor       *string `json:"floor"`
	XSum        *int    `json:"x_sum"`
	YSum        *int    `json:"y_sum"`
}
