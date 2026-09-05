package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type ProjectTransportHandler struct {
	Service *service.ProjectTransportService
}

func NewProjectTransportHandler(s *service.ProjectTransportService) *ProjectTransportHandler {
	return &ProjectTransportHandler{Service: s}
}
func (h *ProjectTransportHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	req, ok := decodeRequest[model.ProjectTransportRequest](w, r)
	if !ok {
		return
	}
	item, err, fields := h.Service.Create(r.Context(), userID, nil, req)
	if err != nil {
		writeResourceError(w, err, fields)
		return
	}
	response.JSON(w, http.StatusCreated, "created", item, nil)
}
func (h *ProjectTransportHandler) CheckerCreate(w http.ResponseWriter, r *http.Request) {
	checkerID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	req, ok := decodeRequest[model.ProjectTransportRequest](w, r)
	if !ok {
		return
	}
	item, err, fields := h.Service.Create(r.Context(), checkerID, &checkerID, req)
	if err != nil {
		writeResourceError(w, err, fields)
		return
	}
	response.JSON(w, http.StatusCreated, "created", item, nil)
}
func (h *ProjectTransportHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.ProjectTransportRequest, model.ProjectTransport](w, r, h.Service.Update)
}
func (h *ProjectTransportHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *ProjectTransportHandler) Get(w http.ResponseWriter, r *http.Request) {
	if id := requestID(r); id > 0 {
		item, err := h.Service.GetByID(r.Context(), id, nil)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}
	opts := parseListOptions(r)
	from, to, err := parseDateRange(r, false)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"date": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	items, err := h.Service.List(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.ProjectTransport]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
func (h *ProjectTransportHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	h.operationalGet(w, r, false, true)
}
func (h *ProjectTransportHandler) DriverGet(w http.ResponseWriter, r *http.Request) {
	h.operationalGet(w, r, true, false)
}
func (h *ProjectTransportHandler) operationalGet(w http.ResponseWriter, r *http.Request, defaultToday, checker bool) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	opts := parseListOptions(r)
	id := requestID(r)
	if id > 0 {
		opts.TransportID = &id
		opts.Limit = 1
		opts.Offset = 0
	}
	if checker {
		opts.CheckerID = &userID
		opts.DriverID = nil
	} else {
		opts.DriverID = &userID
		opts.CheckerID = nil
	}
	from, to, err := parseDateRange(r, defaultToday && id <= 0)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid date filter", nil, map[string]string{"date": "must use YYYY-MM-DD"})
		return
	}
	opts.DateFrom, opts.DateTo = from, to
	items, err := h.Service.ListOperational(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	if id > 0 {
		if len(items) == 0 {
			response.JSON(w, http.StatusNotFound, "transport not found", nil, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", items[0], nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.DriverProjectTransport]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
