package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ReportHandler struct {
	Service *service.ReportService
	Cfg     *config.Config
}

func NewReportHandler(s *service.ReportService, c *config.Config) *ReportHandler {
	return &ReportHandler{Service: s, Cfg: c}
}

func (h *ReportHandler) Get(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	item, err := h.Service.Get(r.Context(), nil, q.Get("period"))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", item, nil)
}
