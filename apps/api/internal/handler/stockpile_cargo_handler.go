package handler

import (
	"net/http"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
)

type StockpileCargoHandler struct {
	Service *service.StockpileCargoService
}

func NewStockpileCargoHandler(s *service.StockpileCargoService, _ *config.Config) *StockpileCargoHandler {
	return &StockpileCargoHandler{Service: s}
}
func (h *StockpileCargoHandler) Create(w http.ResponseWriter, r *http.Request) {
	handleCreate[model.StockpileCargoRequest, model.StockpileCargo](w, r, h.Service.Create)
}
func (h *StockpileCargoHandler) Update(w http.ResponseWriter, r *http.Request) {
	handleUpdate[model.StockpileCargoRequest, model.StockpileCargo](w, r, h.Service.Update)
}
func (h *StockpileCargoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, h.Service.Delete)
}
func (h *StockpileCargoHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleGet[model.StockpileCargo](w, r, h.Service.GetByID, h.Service.List)
}
