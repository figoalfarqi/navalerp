package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type ProjectHandler struct{ Service *service.ProjectService }

func NewProjectHandler(s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{Service: s}
}
func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.ProjectRequest, model.Project](w, r, h.Service.Create)
}
func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectRequest, model.Project](w, r, h.Service.Update)
}
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.get(w, r, nil)
}
func (h *ProjectHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	if id := requestID(r); id > 0 {
		item, err := h.Service.GetByID(r.Context(), id, &checkerID)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", model.NewCheckerProject(*item), nil)
		return
	}
	opts := parseListOptions(r)
	opts.CheckerID = &checkerID
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	checkerProjects := make([]model.CheckerProject, 0, len(items))
	for _, item := range items {
		checkerProjects = append(checkerProjects, model.NewCheckerProject(item))
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.CheckerProject]{
		Items: checkerProjects, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
func (h *ProjectHandler) get(w http.ResponseWriter, r *http.Request, checkerID *int) {
	if id := requestID(r); id > 0 {
		item, err := h.Service.GetByID(r.Context(), id, checkerID)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}
	opts := parseListOptions(r)
	opts.CheckerID = checkerID
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.Project]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
