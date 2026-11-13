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

type warehouseHandler struct {
	warehouseService service.WarehouseServiceInterface
}

func NewWarehouseHandler(warehouseService service.WarehouseServiceInterface) WarehouseHandlerInterface {
	return &warehouseHandler{
		warehouseService: warehouseService,
	}
}

// WarehouseCreate handles create warehouse
// @Summary create warehouse
// @Description create warehouse
// @Tags warehouse
// @Produce json
// @Param	WarehouseRequest body  request.WarehouseRequest true "Create Warehouse Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /warehouse [post]
func (h *warehouseHandler) WarehouseCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.WarehouseDataResponse
	var req request.WarehouseRequest

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

	data, status, err := h.warehouseService.WarehouseCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// WarehouseGetAll handles list warehouse
// @Summary list warehouse
// @Description list all warehouse
// @Tags warehouse
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /warehouse [get]
func (h *warehouseHandler) WarehouseGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.WarehouseListResponse

	data, status, err := h.warehouseService.WarehouseGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// WarehouseGetByID handles get warehouse by id
// @Summary get warehouse by id
// @Description get warehouse by id
// @Tags warehouse
// @Produce json
// @Param	id	path	int	true "Warehouse ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /warehouse/{id} [get]
func (h *warehouseHandler) WarehouseGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.WarehouseDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.warehouseService.WarehouseGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// WarehouseUpdate handles update warehouse by id
// @Summary update warehouse by id
// @Description update warehouse by id
// @Tags warehouse
// @Produce json
// @Param	id	path	int	true "Warehouse ID"
// @Param	WarehouseRequest body  request.WarehouseRequest true "Update Warehouse Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /warehouse/{id} [put]
func (h *warehouseHandler) WarehouseUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.WarehouseResponse
	var req request.WarehouseRequest

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

	status, err := h.warehouseService.WarehouseUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// WarehouseDelete handles delete warehouse by id
// @Summary delete warehouse by id
// @Description delete warehouse by id
// @Tags warehouse
// @Produce json
// @Param	id	path	int	true "Warehouse ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /warehouse/{id} [delete]
func (h *warehouseHandler) WarehouseDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.WarehouseResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.warehouseService.WarehouseDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
