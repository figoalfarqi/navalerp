package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ProjectTruckAssignmentHandler struct {
	Service *service.ProjectTruckAssignmentService
}

func NewProjectTruckAssignmentHandler(s *service.ProjectTruckAssignmentService) *ProjectTruckAssignmentHandler {
	return &ProjectTruckAssignmentHandler{Service: s}
}
func (h *ProjectTruckAssignmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.ProjectTruckAssignmentRequest, model.ProjectTruckAssignment](w, r, h.Service.Create)
}
func (h *ProjectTruckAssignmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectTruckAssignmentRequest, model.ProjectTruckAssignment](w, r, h.Service.Update)
}
func (h *ProjectTruckAssignmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectTruckAssignmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.ProjectTruckAssignment](w, r, h.Service.GetByID, h.Service.List)
}
func (h *ProjectTruckAssignmentHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	opts := parseListOptions(r)
	opts.CheckerID = &checkerID
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectTruckAssignment]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
