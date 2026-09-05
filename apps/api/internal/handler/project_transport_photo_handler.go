package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ProjectTransportPhotoHandler struct {
	Service *service.ProjectTransportPhotoService
}

func NewProjectTransportPhotoHandler(s *service.ProjectTransportPhotoService) *ProjectTransportPhotoHandler {
	return &ProjectTransportPhotoHandler{Service: s}
}
func (h *ProjectTransportPhotoHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.create(w, r, false)
}
func (h *ProjectTransportPhotoHandler) CheckerCreate(w http.ResponseWriter, r *http.Request) {
	h.create(w, r, true)
}
func (h *ProjectTransportPhotoHandler) create(w http.ResponseWriter, r *http.Request, checker bool) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	req, ok := decodeRequest[model.ProjectTransportPhotoRequest](w, r)
	if !ok {
		return
	}
	var checkerID *int
	if checker {
		checkerID = &userID
	}
	item, err, fields := h.Service.Create(r.Context(), userID, checkerID, req)
	if err != nil {
		writeResourceError(w, err, fields)
		return
	}
	response.JSON(w, http.StatusCreated, "created", item, nil)
}
func (h *ProjectTransportPhotoHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectTransportPhotoRequest, model.ProjectTransportPhoto](w, r, h.Service.Update)
}
func (h *ProjectTransportPhotoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectTransportPhotoHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.get(w, r, nil)
}
func (h *ProjectTransportPhotoHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	h.get(w, r, &checkerID)
}
func (h *ProjectTransportPhotoHandler) get(w http.ResponseWriter, r *http.Request, checkerID *int) {
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
	from, to, dateErr := parseDateRange(r, false)
	if dateErr != nil {
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"created_at": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	statusID := intPointer(queryValues(r).Get("project_transport_status_id"))
	items, err := h.Service.List(r.Context(), opts, statusID)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectTransportPhoto]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
