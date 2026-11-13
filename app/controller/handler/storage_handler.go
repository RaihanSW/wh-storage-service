package handler

import (
	"encoding/json"
	"net/http"
	httpresponse "wh-storage-service/app/controller/handler/httpResponse"
	"wh-storage-service/app/model/request"
	"wh-storage-service/app/model/response"
	"wh-storage-service/app/service"
	"wh-storage-service/util"

	"github.com/go-chi/chi/v5"
)

type storageHandler struct {
	storageService service.StorageServiceInterface
}

func NewStorageHandler(storageService service.StorageServiceInterface) StorageHandlerInterface {
	return &storageHandler{
		storageService: storageService,
	}
}

// StorageCreate handles create storage
// @Summary create storage
// @Description create storage
// @Tags storage
// @Produce json
// @Param	StorageRequest body  request.StorageRequest true "Create Storage Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /storage [post]
func (h *storageHandler) StorageCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageDataResponse
	var req request.StorageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	if status, err := util.ValidateRequest(r.Context(), req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	data, status, err := h.storageService.StorageCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageGetAll handles list storage
// @Summary list storage
// @Description list all storage
// @Tags storage
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /storage [get]
func (h *storageHandler) StorageGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageListResponse

	data, status, err := h.storageService.StorageGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageGetByID handles get storage by uid
// @Summary get storage by uid
// @Description get storage by uid
// @Tags storage
// @Produce json
// @Param	uid	path	string	true "Storage UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/{uid} [get]
func (h *storageHandler) StorageGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageDataResponse

	uid := chi.URLParam(r, "uid")

	data, status, err := h.storageService.StorageGetByID(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageUpdate handles update storage by uid
// @Summary update storage by uid
// @Description update storage by uid
// @Tags storage
// @Produce json
// @Param	uid	path	string	true "Storage UID"
// @Param	StorageRequest body  request.StorageRequest true "Update Storage Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/{uid} [put]
func (h *storageHandler) StorageUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageResponse
	var req request.StorageRequest

	uid := chi.URLParam(r, "uid")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	if status, err := util.ValidateRequest(r.Context(), req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	status, err := h.storageService.StorageUpdate(r.Context(), uid, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// StorageDelete handles delete storage by uid
// @Summary delete storage by uid
// @Description delete storage by uid
// @Tags storage
// @Produce json
// @Param	uid	path	string	true "Storage UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/{uid} [delete]
func (h *storageHandler) StorageDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageResponse

	uid := chi.URLParam(r, "uid")

	status, err := h.storageService.StorageDelete(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
