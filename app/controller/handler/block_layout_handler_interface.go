package handler

import "net/http"

type BlockLayoutHandlerInterface interface {
	BlockLayoutCreate(w http.ResponseWriter, r *http.Request)
	BlockLayoutGetAll(w http.ResponseWriter, r *http.Request)
	BlockLayoutGetByID(w http.ResponseWriter, r *http.Request)
	BlockLayoutUpdate(w http.ResponseWriter, r *http.Request)
	BlockLayoutDelete(w http.ResponseWriter, r *http.Request)
}
