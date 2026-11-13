package handler

import "net/http"

type WarehouseHandlerInterface interface {
	WarehouseCreate(w http.ResponseWriter, r *http.Request)
	WarehouseGetAll(w http.ResponseWriter, r *http.Request)
	WarehouseGetByID(w http.ResponseWriter, r *http.Request)
	WarehouseUpdate(w http.ResponseWriter, r *http.Request)
	WarehouseDelete(w http.ResponseWriter, r *http.Request)
}
