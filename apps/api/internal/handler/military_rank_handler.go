package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type MilitaryRankHandler struct {
	Svc *service.MilitaryRankService
	Cfg *config.Config
}

func NewMilitaryRankHandler(s *service.MilitaryRankService, c *config.Config) *MilitaryRankHandler {
	return &MilitaryRankHandler{Svc: s, Cfg: c}
}

func (h *MilitaryRankHandler) Create(w http.ResponseWriter, r *http.Request) {
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.MilitaryRankRequest
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
	response.JSON(w, http.StatusCreated, "military_rank created", item, nil)
}

func (h *MilitaryRankHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id != "" {
		item, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "not found", nil, nil)
			return
		}
		response.JSON(w, http.StatusOK, "success", item, nil)
		return
	}

	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 { page = 1 }
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 { limit = 10 }
	offset := (page - 1) * limit
	search := strings.TrimSpace(query.Get("search"))

	opts := model.ListOptions{Limit: limit, Offset: offset, Search: search}
	items, total, err := h.Svc.List(r.Context(), opts)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "success", map[string]any{
		"items": items,
		"total": total,
		"page": page,
		"limit": limit,
	}, nil)
}

func (h *MilitaryRankHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.JSON(w, http.StatusBadRequest, "id is required", nil, nil)
		return
	}
	loginID, _ := r.Context().Value(middleware.CtxUserID).(string)
	var req model.MilitaryRankRequest
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
	response.JSON(w, http.StatusOK, "military_rank updated", item, nil)
}

func (h *MilitaryRankHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	response.JSON(w, http.StatusOK, "military_rank deleted", nil, nil)
}
