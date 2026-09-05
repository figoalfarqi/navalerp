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

type TruckMerkService struct {
	Repo *repository.TruckMerkRepository
}

func NewTruckMerkService(r *repository.TruckMerkRepository) *TruckMerkService {
	return &TruckMerkService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimTruckMerk(req *model.TruckMerkRequest) {
	validation.TrimStrings(
		&req.TruckMerkName,
	)
}

// ==================================================
// Create
// ==================================================
func (s *TruckMerkService) Create(ctx context.Context, loginID int, req *model.TruckMerkRequest) (*model.TruckMerkResponse, error, map[string]string) {
	trimTruckMerk(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.TruckMerk{
		TruckMerkName: req.TruckMerkName,
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
func (s *TruckMerkService) Update(ctx context.Context, loginID, id int, req *model.TruckMerkRequest) (*model.TruckMerkResponse, error, map[string]string) {
	trimTruckMerk(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.TruckMerk{
		TruckMerkName: req.TruckMerkName,
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
func (s *TruckMerkService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *TruckMerkService) GetByID(ctx context.Context, id int) (*model.TruckMerkResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTruckMerkResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *TruckMerkService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.TruckMerkResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.TruckMerkResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toTruckMerkResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toTruckMerkResp(m *model.TruckMerk) *model.TruckMerkResponse {
	if m == nil {
		return nil
	}

	Trucks := make([]model.TruckResponse, 0)
	if m.Trucks != nil {
		for _, trk := range m.Trucks {
			Trucks = append(Trucks, *toTruckResp(&trk))
		}
	}
	return &model.TruckMerkResponse{
		TruckMerkID:   m.TruckMerkID,
		TruckMerkName: m.TruckMerkName,
		IsActive:      m.IsActive,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     m.DeletedAt,
		Trucks:        Trucks,
	}
}
