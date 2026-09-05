package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
)

type ProjectTransportStatusHandler struct {
	Service *service.ProjectTransportStatusService
}

func NewProjectTransportStatusHandler(s *service.ProjectTransportStatusService) *ProjectTransportStatusHandler {
	return &ProjectTransportStatusHandler{Service: s}
}
func (h *ProjectTransportStatusHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.create(w, r, false)
}
func (h *ProjectTransportStatusHandler) CheckerCreate(w http.ResponseWriter, r *http.Request) {
	h.create(w, r, true)
}
func (h *ProjectTransportStatusHandler) create(w http.ResponseWriter, r *http.Request, checker bool) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	req, ok := decodeRequest[model.ProjectTransportStatusRequest](w, r)
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
func (h *ProjectTransportStatusHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectTransportStatusRequest, model.ProjectTransportStatus](w, r, h.Service.Update)
}
func (h *ProjectTransportStatusHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectTransportStatusHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.get(w, r, nil)
}
func (h *ProjectTransportStatusHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	h.get(w, r, &checkerID)
}
func (h *ProjectTransportStatusHandler) get(w http.ResponseWriter, r *http.Request, checkerID *int) {
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
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"status_time": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	transportID := intPointer(queryValues(r).Get("project_transport_id"))
	items, err := h.Service.List(r.Context(), opts, transportID)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectTransportStatus]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
