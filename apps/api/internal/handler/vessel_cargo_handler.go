package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
)

type VesselCargoHandler struct{ Service *service.VesselCargoService }

func NewVesselCargoHandler(s *service.VesselCargoService) *VesselCargoHandler {
	return &VesselCargoHandler{Service: s}
}
func (h *VesselCargoHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.VesselCargoRequest, model.VesselCargo](w, r, h.Service.Create)
}
func (h *VesselCargoHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.VesselCargoRequest, model.VesselCargo](w, r, h.Service.Update)
}
func (h *VesselCargoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *VesselCargoHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.VesselCargo](w, r, h.Service.GetByID, h.Service.List)
}
