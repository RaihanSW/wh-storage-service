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

type blockLayoutHandler struct {
	blockLayoutService service.BlockLayoutServiceInterface
}

func NewBlockLayoutHandler(blockLayoutService service.BlockLayoutServiceInterface) BlockLayoutHandlerInterface {
	return &blockLayoutHandler{
		blockLayoutService: blockLayoutService,
	}
}

// BlockLayoutCreate handles create block layout
// @Summary create block layout
// @Description create block layout
// @Tags block-layout
// @Produce json
// @Param	BlockLayoutRequest body  request.BlockLayoutRequest true "Create Block Layout Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /block-layout [post]
func (h *blockLayoutHandler) BlockLayoutCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.BlockLayoutDataResponse
	var req request.BlockLayoutRequest

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

	data, status, err := h.blockLayoutService.BlockLayoutCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// BlockLayoutGetAll handles list block layouts
// @Summary list block layouts
// @Description list all block layouts
// @Tags block-layout
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /block-layout [get]
func (h *blockLayoutHandler) BlockLayoutGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.BlockLayoutListResponse

	data, status, err := h.blockLayoutService.BlockLayoutGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// BlockLayoutGetByID handles get block layout by id
// @Summary get block layout by id
// @Description get block layout by id
// @Tags block-layout
// @Produce json
// @Param	id	path	int	true "Block Layout ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /block-layout/{id} [get]
func (h *blockLayoutHandler) BlockLayoutGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.BlockLayoutDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.blockLayoutService.BlockLayoutGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// BlockLayoutUpdate handles update block layout by id
// @Summary update block layout by id
// @Description update block layout by id
// @Tags block-layout
// @Produce json
// @Param	id	path	int	true "Block Layout ID"
// @Param	BlockLayoutRequest body  request.BlockLayoutRequest true "Update Block Layout Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /block-layout/{id} [put]
func (h *blockLayoutHandler) BlockLayoutUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.BlockLayoutResponse
	var req request.BlockLayoutRequest

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

	status, err := h.blockLayoutService.BlockLayoutUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// BlockLayoutDelete handles delete block layout by id
// @Summary delete block layout by id
// @Description delete block layout by id
// @Tags block-layout
// @Produce json
// @Param	id	path	int	true "Block Layout ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /block-layout/{id} [delete]
func (h *blockLayoutHandler) BlockLayoutDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.BlockLayoutResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.blockLayoutService.BlockLayoutDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
