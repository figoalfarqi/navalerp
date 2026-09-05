package handler

import (
	"net/http"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
)

type VesselHandler struct{ Service *service.VesselService }

func NewVesselHandler(s *service.VesselService) *VesselHandler {
	return &VesselHandler{Service: s}
}
func (h *VesselHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.VesselRequest, model.Vessel](w, r, h.Service.Create)
}
func (h *VesselHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.VesselRequest, model.Vessel](w, r, h.Service.Update)
}
func (h *VesselHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *VesselHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.Vessel](w, r, h.Service.GetByID, h.Service.List)
}
