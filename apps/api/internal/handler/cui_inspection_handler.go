package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type CuiInspectionHandler struct {
	Svc *service.CuiInspectionService
	Cfg *config.Config
}

func NewCuiInspectionHandler(svc *service.CuiInspectionService, cfg *config.Config) *CuiInspectionHandler {
	return &CuiInspectionHandler{Svc: svc, Cfg: cfg}
}

func (h *CuiInspectionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id != "" {
		item, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "cui_inspection not found", nil, nil)
			return
		}
		response.JSON(w, http.StatusOK, "cui_inspection retrieved", item, nil)
		return
	}

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	search := q.Get("search")
	if search == "" {
		search = q.Get("q")
	}
	sortBy := q.Get("order_by")
	if sortBy == "" {
		sortBy = q.Get("sort_by")
	}
	order := q.Get("sort")
	if order == "" {
		order = q.Get("order")
	}

	excludedParams := map[string]bool{
		"page":       true,
		"limit":      true,
		"offset":     true,
		"search":     true,
		"q":          true,
		"order_by":   true,
		"sort_by":    true,
		"order":      true,
		"sort":       true,
		"authToken":  true,
		"auth_token": true,
		"_":          true,
	}

	filters := make(map[string]string)
	for k, v := range q {
		if !excludedParams[k] && len(v) > 0 && v[0] != "" {
			filters[k] = v[0]
		}
	}

	opts := model.ListOptions{Limit: limit, Offset: offset, Search: search, SortBy: sortBy, Order: order, Sort: order, Filters: filters}
	items, total, err := h.Svc.List(r.Context(), opts)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "cui_inspection list retrieved", map[string]any{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	}, nil)
}

func (h *CuiInspectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.CuiInspectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	item, err, valErr := h.Svc.Create(r.Context(), loginID, &req)
	if valErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, valErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusCreated, "cui_inspection created", item, nil)
}

func (h *CuiInspectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.JSON(w, http.StatusBadRequest, "id is required", nil, nil)
		return
	}
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.CuiInspectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	item, err, valErr := h.Svc.Update(r.Context(), loginID, id, &req)
	if valErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, valErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "cui_inspection updated", item, nil)
}

func (h *CuiInspectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.JSON(w, http.StatusBadRequest, "id is required", nil, nil)
		return
	}
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	if err := h.Svc.Delete(r.Context(), loginID, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "cui_inspection deleted", nil, nil)
}
