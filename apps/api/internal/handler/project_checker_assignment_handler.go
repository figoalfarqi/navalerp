package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type ProjectCheckerAssignmentHandler struct {
	Service *service.ProjectCheckerAssignmentService
}

func NewProjectCheckerAssignmentHandler(s *service.ProjectCheckerAssignmentService) *ProjectCheckerAssignmentHandler {
	return &ProjectCheckerAssignmentHandler{Service: s}
}
func (h *ProjectCheckerAssignmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.ProjectCheckerAssignmentRequest, model.ProjectCheckerAssignment](w, r, h.Service.Create)
}
func (h *ProjectCheckerAssignmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectCheckerAssignmentRequest, model.ProjectCheckerAssignment](w, r, h.Service.Update)
}
func (h *ProjectCheckerAssignmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectCheckerAssignmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.ProjectCheckerAssignment](w, r, h.Service.GetByID, h.Service.List)
}
func (h *ProjectCheckerAssignmentHandler) CheckerDefault(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	item, err := h.Service.Default(r.Context(), checkerID)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", item, nil)
}
