package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type AppSettingHandler struct{ Service *service.AppSettingService }

func NewAppSettingHandler(s *service.AppSettingService) *AppSettingHandler {
	return &AppSettingHandler{Service: s}
}

func (h *AppSettingHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.AppSettingRequest, model.AppSetting](w, r, h.Service.Create)
}

func (h *AppSettingHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.AppSettingRequest, model.AppSetting](w, r, h.Service.Update)
}

func (h *AppSettingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}

func (h *AppSettingHandler) Get(w http.ResponseWriter, r *http.Request) {
	if id := requestID(r); id > 0 {
		item, err := h.Service.GetByID(r.Context(), id)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}
	opts := parseListOptions(r)
	items, err := h.Service.List(r.Context(), opts, false)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.AppSetting]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}

func (h *AppSettingHandler) CheckerGet(w http.ResponseWriter, r *http.Request) {
	opts := parseListOptions(r)
	items, err := h.Service.List(r.Context(), opts, true)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[model.AppSetting]{
		Items: items, Limit: opts.Limit, Offset: opts.Offset,
	}, nil)
}
