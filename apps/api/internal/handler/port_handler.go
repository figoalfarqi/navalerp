package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
)

type PortHandler struct{ Service *service.PortService }

func NewPortHandler(s *service.PortService) *PortHandler { return &PortHandler{Service: s} }
func (h *PortHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.PortRequest, model.Port](w, r, h.Service.Create)
}
func (h *PortHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.PortRequest, model.Port](w, r, h.Service.Update)
}
func (h *PortHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *PortHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.Port](w, r, h.Service.GetByID, h.Service.List)
}
