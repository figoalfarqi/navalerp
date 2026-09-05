package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type ProjectRouteHandler struct{ Service *service.ProjectRouteService }

func NewProjectRouteHandler(s *service.ProjectRouteService) *ProjectRouteHandler {
	return &ProjectRouteHandler{Service: s}
}
func (h *ProjectRouteHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.ProjectRouteRequest, model.ProjectRoute](w, r, h.Service.Create)
}
func (h *ProjectRouteHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectRouteRequest, model.ProjectRoute](w, r, h.Service.Update)
}
func (h *ProjectRouteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectRouteHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.get(w, r, nil)
}
func (h *ProjectRouteHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	h.get(w, r, &checkerID)
}
func (h *ProjectRouteHandler) get(w http.ResponseWriter, r *http.Request, checkerID *int) {
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
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectRoute]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
