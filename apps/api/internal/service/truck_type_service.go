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

type TruckTypeService struct {
	Repo *repository.TruckTypeRepository
}

func NewTruckTypeService(r *repository.TruckTypeRepository) *TruckTypeService {
	return &TruckTypeService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimTruckType(req *model.TruckTypeRequest) {
	validation.TrimStrings(&req.TruckTypeName, req.TruckTypeDescription)
}

// ==================================================
// Create
// ==================================================
func (s *TruckTypeService) Create(ctx context.Context, loginID int, req *model.TruckTypeRequest) (*model.TruckTypeResponse, error, map[string]string) {
	trimTruckType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	// default is_active = 1
	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	ft := &model.TruckType{
		TruckTypeName:        req.TruckTypeName,
		TruckTypeDescription: req.TruckTypeDescription,
		TruckBoxLength:       req.TruckBoxLength,
		TruckBoxWidth:        req.TruckBoxWidth,
		TruckBoxHeight:       req.TruckBoxHeight,
		TruckCapacity:        req.TruckCapacity,
		IsActive:             isActive,
		CreatedBy:            loginID,
		UpdatedBy:            loginID,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	id, err := s.Repo.Create(ctx, ft)
	if err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Update
// ==================================================
func (s *TruckTypeService) Update(ctx context.Context, loginID, id int, req *model.TruckTypeRequest) (*model.TruckTypeResponse, error, map[string]string) {
	trimTruckType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	ft := &model.TruckType{
		TruckTypeName:        req.TruckTypeName,
		TruckTypeDescription: req.TruckTypeDescription,
		TruckBoxLength:       req.TruckBoxLength,
		TruckBoxWidth:        req.TruckBoxWidth,
		TruckBoxHeight:       req.TruckBoxHeight,
		TruckCapacity:        req.TruckCapacity,
		UpdatedBy:            loginID,
		UpdatedAt:            time.Now(),
	}

	// update is_active hanya jika user mengirimkan nilainya
	if req.IsActive != nil {
		ft.IsActive = *req.IsActive
	} else {
		ft.IsActive = -1 // signal: jangan update column is_active
	}

	if err := s.Repo.Update(ctx, id, ft); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Delete (Soft Delete)
// ==================================================
func (s *TruckTypeService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *TruckTypeService) GetByID(ctx context.Context, id int) (*model.TruckTypeResponse, error) {
	ft, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTruckTypeResp(ft), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *TruckTypeService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.TruckTypeResponse, error) {
	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.TruckTypeResponse, 0, len(items))
	for _, ft := range items {
		resp = append(resp, *toTruckTypeResp(&ft))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toTruckTypeResp(ft *model.TruckType) *model.TruckTypeResponse {
	if ft == nil {
		return nil
	}

	return &model.TruckTypeResponse{
		TruckTypeID:          ft.TruckTypeID,
		TruckTypeName:        ft.TruckTypeName,
		TruckTypeDescription: ft.TruckTypeDescription,
		TruckBoxLength:       ft.TruckBoxLength,
		TruckBoxWidth:        ft.TruckBoxWidth,
		TruckBoxHeight:       ft.TruckBoxHeight,
		TruckCapacity:        ft.TruckCapacity,
		IsActive:             ft.IsActive,
		CreatedAt:            ft.CreatedAt,
		UpdatedAt:            ft.UpdatedAt,
		DeletedAt:            ft.DeletedAt,
	}
}
