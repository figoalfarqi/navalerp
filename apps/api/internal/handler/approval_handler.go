package handler

import (
	"encoding/json"
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ApprovalHandler struct {
	Svc *service.ApprovalService
	Cfg *config.Config
}

func NewApprovalHandler(svc *service.ApprovalService, cfg *config.Config) *ApprovalHandler {
	return &ApprovalHandler{Svc: svc, Cfg: cfg}
}

func (h *ApprovalHandler) Process(w http.ResponseWriter, r *http.Request) {
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	if loginID == "" {
		loginID = "SYSTEM_ADMIN"
	}

	var req model.ApprovalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid JSON body", nil, map[string]string{"error": err.Error()})
		return
	}

	res, err := h.Svc.ProcessApproval(r.Context(), loginID, &req)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	response.JSON(w, http.StatusOK, "Approval processed successfully", res, nil)
}

func (h *ApprovalHandler) ListPending(w http.ResponseWriter, r *http.Request) {
	items, err := h.Svc.GetPendingApprovals(r.Context())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.JSON(w, http.StatusOK, "Pending approvals retrieved", items, nil)
}
