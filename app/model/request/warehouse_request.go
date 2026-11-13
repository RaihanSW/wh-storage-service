package request

type WarehouseRequest struct {
	Name    string  `json:"name" validate:"required"`
	Code    string  `json:"code" validate:"required"`
	Address *string `json:"address"`
	TtlArea *int    `json:"ttl_area"`
}
