package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/app/middleware"
	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/service"
	"github.com/figoalfarqi/apipml/pkg/response"
	"github.com/jackc/pgx/v5"
)

type TruckHandler struct {
	Svc *service.TruckService
	Cfg *config.Config
}

func NewTruckHandler(s *service.TruckService, c *config.Config) *TruckHandler {
	return &TruckHandler{Svc: s, Cfg: c}
}

// ==================================================
// POST /api/v1/admin/truck
// ==================================================
func (h *TruckHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	// get user login id
	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	var req model.TruckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}

	item, err, formErr := h.Svc.Create(r.Context(), loginID, &req)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to create truck", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusCreated, "truck created", item, nil)
}

// ==================================================
// PUT /api/v1/admin/truck/{id}
// ==================================================
func (h *TruckHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	// ambil id dari path
	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "truck")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	var req model.TruckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}

	item, err, formErr := h.Svc.Update(r.Context(), loginID, id, &req)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to update truck", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "truck updated", item, nil)
}

// ==================================================
// DELETE /api/v1/admin/truck/{id}
// ==================================================
func (h *TruckHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "truck")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	if err := h.Svc.Delete(r.Context(), loginID, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to delete truck", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "truck deleted", map[string]int{
		"truck_id": id,
	}, nil)
}

// ==================================================
// GET /api/v1/admin/truck OR /{id}
// ==================================================
func (h *TruckHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "truck")

	// ==================================================
	// GET BY ID
	// ==================================================
	if id > 0 {
		item, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "truck not found", nil, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "failed to fetch truck", nil, map[string]string{
				"error": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}

	// ==================================================
	// LIST (cursor + filter)
	// ==================================================
	queryParams := r.URL.Query()
	limit := helper.AtoiSafeDefault(queryParams.Get("limit"), 10)
	cursorData, err := helper.ParseCursorParams(r.URL.Query())
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid cursor", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// allowed filters
	baseAllowed := []string{
		"license_plate",
		"truck_type_id",
		"truck_merk_id",
		"driver_id",
		"ownership_status_id",
		"production_year",
		"number_of_tires",
		"created_by",
		"updated_by",
		"is_active",
	}

	filterOnly := []string{
		"created_at_after",
		"created_at_before",
		"updated_at_after",
		"updated_at_before",
	}

	sortOnly := []string{
		"created_at",
		"updated_at",
	}

	allowedFilters := append(baseAllowed, filterOnly...)
	allowedSorts := append(baseAllowed, sortOnly...)

	filters := helper.ParseQueryFilter(queryParams, allowedFilters)
	orderBy, sort := helper.ParseQuerySort(queryParams, allowedSorts)

	var cursorValue interface{} = nil
	var cursorKey *int = nil
	if cursorData != nil {
		cursorValue = cursorData.Value
		cursorKey = cursorData.Key
	}

	items, err := h.Svc.List(r.Context(), cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to list trucks", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "ok", map[string]interface{}{
		"items":        items,
		"limit":        limit,
		"cursor_value": cursorValue,
		"cursor_key":   cursorKey,
		"filters":      filters,
		"order_by":     orderBy,
		"sort":         sort,
	}, nil)
}

// ==================================================
// GET /api/v1/admin/truck_with_delivery_order
// ==================================================
// func (h *TruckHandler) GetWithDeliveryOrder(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
// 		return
// 	}

// 	// ==================================================
// 	// LIST (cursor + filter)
// 	// ==================================================
// 	queryParams := r.URL.Query()
// 	limit := helper.AtoiSafeDefault(queryParams.Get("limit"), 10)
// 	cursorData, err := helper.ParseCursorParams(r.URL.Query())
// 	if err != nil {
// 		response.JSON(w, http.StatusBadRequest, "invalid cursor", nil, map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// allowed filters
// 	baseAllowed := []string{
// 		"license_plate",
// 		"truck_type_id",
// 		"truck_merk_id",
// 		"driver_id",
// 		"ownership_status_id",
// 		"production_year",
// 		"number_of_tires",
// 		"last_status_type_id",
// 		"delivery_order_number",
// 	}

// 	filterOnly := []string{
// 		"completed_at_after",
// 		"completed_at_before",
// 		"province_id",
// 		"city_id",
// 		"district_id",
// 	}

// 	sortOnly := []string{
// 		"completed_at",
// 	}

// 	allowedFilters := append(baseAllowed, filterOnly...)
// 	allowedSorts := append(baseAllowed, sortOnly...)

// 	filters := helper.ParseQueryFilter(queryParams, allowedFilters)
// 	orderBy, sort := helper.ParseQuerySort(queryParams, allowedSorts)

// 	var cursorValue interface{} = nil
// 	var cursorKey *int = nil
// 	if cursorData != nil {
// 		cursorValue = cursorData.Value
// 		cursorKey = cursorData.Key
// 	}

// 	items, err := h.Svc.ListWithDeliveryOrder(r.Context(), cursorValue, cursorKey, limit, filters, orderBy, sort)
// 	if err != nil {
// 		response.JSON(w, http.StatusInternalServerError, "failed to list trucks", nil, map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	response.JSON(w, http.StatusOK, "ok", map[string]interface{}{
// 		"items":        items,
// 		"limit":        limit,
// 		"cursor_value": cursorValue,
// 		"cursor_key":   cursorKey,
// 		"filters":      filters,
// 		"order_by":     orderBy,
// 		"sort":         sort,
// 	}, nil)
// }
