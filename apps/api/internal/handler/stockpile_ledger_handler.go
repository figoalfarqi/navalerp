package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type StockpileLedgerHandler struct {
	Service *service.StockpileLedgerService
}

func NewStockpileLedgerHandler(s *service.StockpileLedgerService, _ *config.Config) *StockpileLedgerHandler {
	return &StockpileLedgerHandler{Service: s}
}
func (h *StockpileLedgerHandler) Get(w http.ResponseWriter, r *http.Request) {
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
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"created_at": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.StockpileLedger]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
