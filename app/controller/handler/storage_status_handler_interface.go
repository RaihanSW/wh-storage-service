package handler

import "net/http"

type StorageStatusHandlerInterface interface {
	StorageStatusCreate(w http.ResponseWriter, r *http.Request)
	StorageStatusGetAll(w http.ResponseWriter, r *http.Request)
	StorageStatusGetByID(w http.ResponseWriter, r *http.Request)
	StorageStatusUpdate(w http.ResponseWriter, r *http.Request)
	StorageStatusDelete(w http.ResponseWriter, r *http.Request)
}
