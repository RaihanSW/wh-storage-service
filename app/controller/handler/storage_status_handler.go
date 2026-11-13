package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	httpresponse "wh-storage-service/app/controller/handler/httpResponse"
	"wh-storage-service/app/model/request"
	"wh-storage-service/app/model/response"
	"wh-storage-service/app/service"
	"wh-storage-service/util"

	"github.com/go-chi/chi/v5"
)

type storageStatusHandler struct {
	storageStatusService service.StorageStatusServiceInterface
}

func NewStorageStatusHandler(storageStatusService service.StorageStatusServiceInterface) StorageStatusHandlerInterface {
	return &storageStatusHandler{
		storageStatusService: storageStatusService,
	}
}

// StorageStatusCreate handles create storage status
// @Summary create storage status
// @Description create storage status
// @Tags storage-status
// @Produce json
// @Param	StorageStatusRequest body  request.StorageStatusRequest true "Create Storage Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /storage/status [post]
func (h *storageStatusHandler) StorageStatusCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageStatusDataResponse
	var req request.StorageStatusRequest

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

	data, status, err := h.storageStatusService.StorageStatusCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageStatusGetAll handles list storage statuses
// @Summary list storage statuses
// @Description list all storage statuses
// @Tags storage-status
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /storage/status [get]
func (h *storageStatusHandler) StorageStatusGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageStatusListResponse

	data, status, err := h.storageStatusService.StorageStatusGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageStatusGetByID handles get storage status by id
// @Summary get storage status by id
// @Description get storage status by id
// @Tags storage-status
// @Produce json
// @Param	id	path	int	true "Storage Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/status/{id} [get]
func (h *storageStatusHandler) StorageStatusGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageStatusDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.storageStatusService.StorageStatusGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// StorageStatusUpdate handles update storage status by id
// @Summary update storage status by id
// @Description update storage status by id
// @Tags storage-status
// @Produce json
// @Param	id	path	int	true "Storage Status ID"
// @Param	StorageStatusRequest body  request.StorageStatusRequest true "Update Storage Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/status/{id} [put]
func (h *storageStatusHandler) StorageStatusUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageStatusResponse
	var req request.StorageStatusRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

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

	status, err := h.storageStatusService.StorageStatusUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// StorageStatusDelete handles delete storage status by id
// @Summary delete storage status by id
// @Description delete storage status by id
// @Tags storage-status
// @Produce json
// @Param	id	path	int	true "Storage Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /storage/status/{id} [delete]
func (h *storageStatusHandler) StorageStatusDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.StorageStatusResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.storageStatusService.StorageStatusDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
