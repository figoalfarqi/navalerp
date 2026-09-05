package service

import (
	"context"
	"errors"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
	"github.com/figoalfarqi/navalerp/internal/validation"
)

type StockpileService struct {
	Repo *repository.StockpileRepository
}

func NewStockpileService(r *repository.StockpileRepository) *StockpileService {
	return &StockpileService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimStockpile(req *model.StockpileRequest) {
	validation.TrimStrings(
		&req.StockpileName,
		req.StockpileAddress,
		req.StockpileMapUrl,
	)
}

// ==================================================
// Create
// ==================================================
func (s *StockpileService) Create(ctx context.Context, loginID int, req *model.StockpileRequest) (*model.StockpileResponse, error, map[string]string) {
	trimStockpile(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.Stockpile{
		StockpileName:      req.StockpileName,
		CityID:             req.CityID,
		StockpileAddress:   req.StockpileAddress,
		StockpileLatitude:  req.StockpileLatitude,
		StockpileLongitude: req.StockpileLongitude,
		StockpileMapUrl:    req.StockpileMapUrl,
		IsActive:           isActive,
		CreatedBy:          loginID,
		UpdatedBy:          loginID,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Update
// ==================================================
func (s *StockpileService) Update(ctx context.Context, loginID, id int, req *model.StockpileRequest) (*model.StockpileResponse, error, map[string]string) {
	trimStockpile(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.Stockpile{
		StockpileName:      req.StockpileName,
		CityID:             req.CityID,
		StockpileAddress:   req.StockpileAddress,
		StockpileLatitude:  req.StockpileLatitude,
		StockpileLongitude: req.StockpileLongitude,
		StockpileMapUrl:    req.StockpileMapUrl,
		UpdatedBy:          loginID,
		UpdatedAt:          time.Now(),
	}

	// is_active opsional
	if req.IsActive != nil {
		m.IsActive = *req.IsActive
	} else {
		m.IsActive = -1 // signal ke repo: jangan update is_active
	}

	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Delete (Soft Delete)
// ==================================================
func (s *StockpileService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *StockpileService) GetByID(ctx context.Context, id int) (*model.StockpileResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toStockpileResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *StockpileService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.StockpileResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.StockpileResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toStockpileResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toStockpileResp(m *model.Stockpile) *model.StockpileResponse {
	if m == nil {
		return nil
	}

	var cityResp *model.CityResponse
	if m.City != nil {
		cityResp = toCityResp(m.City)
	}
	StockpileCargos := make([]model.StockpileCargoResponse, 0)
	if m.StockpileCargos != nil {
		for _, stockpileCargo := range m.StockpileCargos {
			StockpileCargos = append(StockpileCargos, *toStockpileCargoResp(&stockpileCargo))
		}
	}
	resp := &model.StockpileResponse{
		StockpileID:        m.StockpileID,
		StockpileName:      m.StockpileName,
		CityID:             m.CityID,
		StockpileAddress:   m.StockpileAddress,
		StockpileLatitude:  m.StockpileLatitude,
		StockpileLongitude: m.StockpileLongitude,
		StockpileMapUrl:    m.StockpileMapUrl,
		IsActive:           m.IsActive,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		DeletedAt:          m.DeletedAt,
		City:               cityResp,
		StockpileCargos:    StockpileCargos,
	}

	return resp
}
