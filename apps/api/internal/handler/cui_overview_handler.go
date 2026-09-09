package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type CuiOverviewHandler struct {
	Svc *service.CuiOverviewService
	Cfg *config.Config
}

func NewCuiOverviewHandler(svc *service.CuiOverviewService, cfg *config.Config) *CuiOverviewHandler {
	return &CuiOverviewHandler{Svc: svc, Cfg: cfg}
}

func (h *CuiOverviewHandler) Get(w http.ResponseWriter, r *http.Request) {
	data, err := h.Svc.GetOverview(r.Context())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.JSON(w, http.StatusOK, "CUI national overview retrieved", data, nil)
}
