package entity

import "time"

type Storage struct {
	UID         string    `json:"uid"`
	Name        string    `json:"name"`
	GoodsUID    *string   `json:"goods_uid,omitempty"`
	WarehouseID int       `json:"warehouse_id"`
	BlockID     int       `json:"block_id"`
	Status      int       `json:"status"`
	XRow        *int      `json:"x_row,omitempty"`
	YRow        *int      `json:"y_row,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
