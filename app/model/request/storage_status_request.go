package request

type StorageStatusRequest struct {
	Name string `json:"name" validate:"required"`
}
