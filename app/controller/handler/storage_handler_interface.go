package handler

import "net/http"

type StorageHandlerInterface interface {
	StorageCreate(w http.ResponseWriter, r *http.Request)
	StorageGetAll(w http.ResponseWriter, r *http.Request)
	StorageGetByID(w http.ResponseWriter, r *http.Request)
	StorageUpdate(w http.ResponseWriter, r *http.Request)
	StorageDelete(w http.ResponseWriter, r *http.Request)
}
