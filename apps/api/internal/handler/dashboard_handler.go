package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type DashboardHandler struct {
	Service *service.DashboardService
	Cfg     *config.Config
}

func NewDashboardHandler(s *service.DashboardService, c *config.Config) *DashboardHandler {
	return &DashboardHandler{Service: s, Cfg: c}
}

func (h *DashboardHandler) Get(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	item, err := h.Service.Get(r.Context(), nil, q.Get("month"))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", item, nil)
}
