package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type ReportHandler struct{ Service *service.ReportService }

func NewReportHandler(s *service.ReportService) *ReportHandler {
	return &ReportHandler{Service: s}
}

func (h *ReportHandler) Get(w http.ResponseWriter, r *http.Request) {
	q := queryValues(r)
	item, err := h.Service.Get(
		r.Context(),
		intPointer(q.Get("project_id")),
		q.Get("period"),
		q.Get("date"),
	)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid report filter", nil, map[string]string{"filter": err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, "ok", item, nil)
}
