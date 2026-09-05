package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type DashboardHandler struct{ Service *service.DashboardService }

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{Service: s}
}

func (h *DashboardHandler) Get(w http.ResponseWriter, r *http.Request) {
	q := queryValues(r)
	item, err := h.Service.Get(r.Context(), intPointer(q.Get("project_id")), q.Get("month"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid dashboard filter", nil, map[string]string{"filter": err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, "ok", item, nil)
}
