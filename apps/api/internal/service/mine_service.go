package service

import (
	"context"
	"errors"
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
	"github.com/figoalfarqi/apipml/internal/validation"
)

type MineService struct {
	Repo *repository.MineRepository
}

func NewMineService(r *repository.MineRepository) *MineService {
	return &MineService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimMine(req *model.MineRequest) {
	validation.TrimStrings(
		&req.MineName,
		req.MineMapUrl,
	)
}

// ==================================================
// Create
// ==================================================
func (s *MineService) Create(ctx context.Context, loginID int, req *model.MineRequest) (*model.MineResponse, error, map[string]string) {
	trimMine(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.Mine{
		MineName:      req.MineName,
		CityID:        req.CityID,
		MineAddress:   req.MineAddress,
		MineLatitude:  req.MineLatitude,
		MineLongitude: req.MineLongitude,
		MineMapUrl:    req.MineMapUrl,
		IsActive:      isActive,
		CreatedBy:     loginID,
		UpdatedBy:     loginID,
		CreatedAt:     now,
		UpdatedAt:     now,
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
func (s *MineService) Update(ctx context.Context, loginID, id int, req *model.MineRequest) (*model.MineResponse, error, map[string]string) {
	trimMine(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.Mine{
		MineName:      req.MineName,
		CityID:        req.CityID,
		MineAddress:   req.MineAddress,
		MineLatitude:  req.MineLatitude,
		MineLongitude: req.MineLongitude,
		MineMapUrl:    req.MineMapUrl,
		UpdatedBy:     loginID,
		UpdatedAt:     time.Now(),
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
func (s *MineService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *MineService) GetByID(ctx context.Context, id int) (*model.MineResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toMineResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *MineService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.MineResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.MineResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toMineResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toMineResp(m *model.Mine) *model.MineResponse {
	if m == nil {
		return nil
	}

	var cityResp *model.CityResponse
	if m.City != nil {
		cityResp = toCityResp(m.City)
	}
	resp := &model.MineResponse{
		MineID:        m.MineID,
		MineName:      m.MineName,
		CityID:        m.CityID,
		MineAddress:   m.MineAddress,
		MineLatitude:  m.MineLatitude,
		MineLongitude: m.MineLongitude,
		MineMapUrl:    m.MineMapUrl,
		IsActive:      m.IsActive,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     m.DeletedAt,
		City:          cityResp,
	}

	return resp
}
