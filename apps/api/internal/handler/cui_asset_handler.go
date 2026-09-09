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

type CuiAssetHandler struct {
	Svc *service.CuiAssetService
	Cfg *config.Config
}

func NewCuiAssetHandler(svc *service.CuiAssetService, cfg *config.Config) *CuiAssetHandler {
	return &CuiAssetHandler{Svc: svc, Cfg: cfg}
}

func (h *CuiAssetHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id != "" {
		item, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "cui_asset not found", nil, nil)
			return
		}
		response.JSON(w, http.StatusOK, "cui_asset retrieved", item, nil)
		return
	}

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 { page = 1 }
	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit < 1 { limit = 10 }
	offset := (page - 1) * limit
	search := q.Get("search")
	sortBy := q.Get("sort_by")
	order := q.Get("order")
	filters := make(map[string]string)
	for k, v := range q {
		if k != "page" && k != "limit" && k != "search" && k != "sort_by" && k != "order" && len(v) > 0 {
			filters[k] = v[0]
		}
	}

	opts := model.ListOptions{Limit: limit, Offset: offset, Search: search, SortBy: sortBy, Order: order, Filters: filters}
	items, total, err := h.Svc.List(r.Context(), opts)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "cui_asset list retrieved", map[string]any{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	}, nil)
}

func (h *CuiAssetHandler) Create(w http.ResponseWriter, r *http.Request) {
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.CuiAssetRequest
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
	response.JSON(w, http.StatusCreated, "cui_asset created", item, nil)
}

func (h *CuiAssetHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.JSON(w, http.StatusBadRequest, "id is required", nil, nil)
		return
	}
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.CuiAssetRequest
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
	response.JSON(w, http.StatusOK, "cui_asset updated", item, nil)
}

func (h *CuiAssetHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	response.JSON(w, http.StatusOK, "cui_asset deleted", nil, nil)
}
