package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ProjectFinancialTransactionHandler struct {
	Service *service.ProjectFinancialTransactionService
}

func NewProjectFinancialTransactionHandler(s *service.ProjectFinancialTransactionService) *ProjectFinancialTransactionHandler {
	return &ProjectFinancialTransactionHandler{Service: s}
}
func (h *ProjectFinancialTransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.ProjectFinancialTransactionRequest, model.ProjectFinancialTransaction](w, r, h.Service.Create)
}
func (h *ProjectFinancialTransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectFinancialTransactionRequest, model.ProjectFinancialTransaction](w, r, h.Service.Update)
}
func (h *ProjectFinancialTransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectFinancialTransactionHandler) Get(w http.ResponseWriter, r *http.Request) {
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
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"transaction_date": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectFinancialTransaction]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
