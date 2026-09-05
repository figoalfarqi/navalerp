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

type FileUploadHandler struct {
	Svc *service.FileUploadService
	Cfg *config.Config
}

func NewFileUploadHandler(s *service.FileUploadService, c *config.Config) *FileUploadHandler {
	return &FileUploadHandler{Svc: s, Cfg: c}
}

// POST /api/v1/allrole/presigned
func (h *FileUploadHandler) GetPresignedURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	var req model.PresignedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}

	idAppUser, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "user not found in context", nil, nil)
		return
	}

	result, err := h.Svc.GeneratePresignedURL(r.Context(), req.Folder, req.Filename, req.Action, idAppUser)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to generate presigned url", nil, nil)
		return
	}

	response.JSON(w, http.StatusCreated, "presigned url generated", result, nil)
}
