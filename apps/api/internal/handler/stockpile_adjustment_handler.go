package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type StockpileAdjustmentHandler struct {
	Service *service.StockpileAdjustmentService
}

func NewStockpileAdjustmentHandler(s *service.StockpileAdjustmentService, _ *config.Config) *StockpileAdjustmentHandler {
	return &StockpileAdjustmentHandler{Service: s}
}
func (h *StockpileAdjustmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.StockpileAdjustmentRequest, model.StockpileAdjustment](w, r, h.Service.Create)
}
func (h *StockpileAdjustmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.StockpileAdjustmentRequest, model.StockpileAdjustment](w, r, h.Service.Update)
}
func (h *StockpileAdjustmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *StockpileAdjustmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	if id := requestID(r); id > 0 {
		item, err := h.Service.GetByID(r.Context(), id)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}
	opts := parseListOptions(r)
	from, to, err := parseDateRange(r, false)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"adjustment_date": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.StockpileAdjustment]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
