package entity

import "time"

type Warehouse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Address   *string   `json:"address,omitempty"`
	TtlArea   *int      `json:"ttl_area,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
