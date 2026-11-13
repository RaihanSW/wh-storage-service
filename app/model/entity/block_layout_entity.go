package entity

import "time"

type BlockLayout struct {
	ID          int       `json:"id"`
	WarehouseID int       `json:"warehouse_id"`
	BlockName   string    `json:"block_name"`
	Floor       *string   `json:"floor,omitempty"`
	XSum        *int      `json:"x_sum,omitempty"`
	YSum        *int      `json:"y_sum,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
