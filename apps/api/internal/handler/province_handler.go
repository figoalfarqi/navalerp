package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
	"github.com/jackc/pgx/v5"
)

type ProvinceHandler struct {
	Svc *service.ProvinceService
	Cfg *config.Config
}

func NewProvinceHandler(s *service.ProvinceService, c *config.Config) *ProvinceHandler {
	return &ProvinceHandler{Svc: s, Cfg: c}
}

// ==================================================
// POST /api/v1/admin/province
// POST /api/v1/driver/province
// ==================================================
func (h *ProvinceHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	var req model.ProvinceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}

	prov, err, formErr := h.Svc.Create(r.Context(), loginID, &req)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to create province", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusCreated, "province created", prov, nil)
}

// ==================================================
// PUT /api/v1/admin/province/{id}
// PUT /api/v1/driver/province/{id}
// ==================================================
func (h *ProvinceHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "province")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{
			"id": "invalid",
		})
		return
	}

	var req model.ProvinceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}

	prov, err, formErr := h.Svc.Update(r.Context(), loginID, id, &req)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to update province", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "province updated", prov, nil)
}

// ==================================================
// DELETE /api/v1/admin/province/{id}
// DELETE /api/v1/driver/province/{id}
// ==================================================
func (h *ProvinceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "province")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{
			"id": "invalid",
		})
		return
	}

	if err := h.Svc.Delete(r.Context(), loginID, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to delete province", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "province deleted", map[string]int{
		"province_id": id,
	}, nil)
}

// ==================================================
// GET /api/v1/admin/province OR /{id}
// GET /api/v1/driver/province OR /{id}
// ==================================================
func (h *ProvinceHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "province")

	// =========================
	// GET by ID
	// =========================
	if id > 0 {
		prov, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "province not found", nil, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "failed to fetch province", nil, map[string]string{
				"error": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, "ok", prov, nil)
		return
	}

	// =========================
	// LIST Province
	// =========================
	queryParams := r.URL.Query()

	limit := helper.AtoiSafeDefault(queryParams.Get("limit"), 10)
	cursorData, err := helper.ParseCursorParams(r.URL.Query())
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid cursor", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	baseAllowed := []string{
		"province_name",
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
		response.JSON(w, http.StatusInternalServerError, "failed to list provinces", nil, map[string]string{
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
