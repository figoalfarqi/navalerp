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

type AppRoleHandler struct {
	Svc *service.AppRoleService
	Cfg *config.Config
}

func NewAppRoleHandler(s *service.AppRoleService, c *config.Config) *AppRoleHandler {
	return &AppRoleHandler{Svc: s, Cfg: c}
}

// ========================================================================
// POST /api/v1/admin/app_role
// ========================================================================
func (h *AppRoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	var req model.AppRoleRequest
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
		response.JSON(w, http.StatusInternalServerError, "failed to create app role", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusCreated, "app role created", item, nil)
}

// ========================================================================
// PUT/PATCH /api/v1/admin/app_role/{id}
// ========================================================================
func (h *AppRoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "app_role")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	var req model.AppRoleRequest
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
		response.JSON(w, http.StatusInternalServerError, "failed to update app role", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "app role updated", item, nil)
}

// ========================================================================
// DELETE /api/v1/admin/app_role/{id}
// ========================================================================
func (h *AppRoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "app_role")
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	if err := h.Svc.Delete(r.Context(), loginID, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to delete app role", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "app role deleted", map[string]int{
		"app_role_id": id,
	}, nil)
}

// ========================================================================
// GET /api/v1/admin/app_role or app_role/{id}
// ========================================================================
func (h *AppRoleHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	_, id := helper.GetIDAndRoleFromPath(r.URL.Path, "app_role")

	// ======================
	// GET BY ID
	// ======================
	if id > 0 {
		item, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "app role not found", nil, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "failed to fetch app role", nil, map[string]string{
				"error": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}

	// ======================
	// LIST (cursor + filters)
	// ======================
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
		"app_role_name",
		"app_role_type_id",
		"app_role_description",
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
		response.JSON(w, http.StatusInternalServerError, "failed to list app roles", nil, map[string]string{
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
