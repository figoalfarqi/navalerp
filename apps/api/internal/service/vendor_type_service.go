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

type VendorTypeService struct {
	Repo *repository.VendorTypeRepository
}

func NewVendorTypeService(r *repository.VendorTypeRepository) *VendorTypeService {
	return &VendorTypeService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimVendorType(req *model.VendorTypeRequest) {
	validation.TrimStrings(&req.VendorTypeName)
}

// ==================================================
// Create
// ==================================================
func (s *VendorTypeService) Create(ctx context.Context, loginID int, req *model.VendorTypeRequest) (*model.VendorTypeResponse, error, map[string]string) {
	trimVendorType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	// default is_active = 1
	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	ft := &model.VendorType{
		VendorTypeName:        req.VendorTypeName,
		VendorTypeDescription: req.VendorTypeDescription,
		IsActive:              isActive,
		CreatedBy:             loginID,
		UpdatedBy:             loginID,
		CreatedAt:             now,
		UpdatedAt:             now,
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
func (s *VendorTypeService) Update(ctx context.Context, loginID, id int, req *model.VendorTypeRequest) (*model.VendorTypeResponse, error, map[string]string) {
	trimVendorType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	ft := &model.VendorType{
		VendorTypeName:        req.VendorTypeName,
		VendorTypeDescription: req.VendorTypeDescription,
		UpdatedBy:             loginID,
		UpdatedAt:             time.Now(),
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
func (s *VendorTypeService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *VendorTypeService) GetByID(ctx context.Context, id int) (*model.VendorTypeResponse, error) {
	ft, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toVendorTypeResp(ft), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *VendorTypeService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.VendorTypeResponse, error) {
	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.VendorTypeResponse, 0, len(items))
	for _, ft := range items {
		resp = append(resp, *toVendorTypeResp(&ft))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toVendorTypeResp(ft *model.VendorType) *model.VendorTypeResponse {
	if ft == nil {
		return nil
	}

	Vendors := make([]model.VendorResponse, 0)
	if ft.Vendors != nil {
		for _, vendor := range ft.Vendors {
			Vendors = append(Vendors, *toVendorResp(&vendor))
		}
	}
	return &model.VendorTypeResponse{
		VendorTypeID:          ft.VendorTypeID,
		VendorTypeName:        ft.VendorTypeName,
		VendorTypeDescription: ft.VendorTypeDescription,
		IsActive:              ft.IsActive,
		CreatedAt:             ft.CreatedAt,
		UpdatedAt:             ft.UpdatedAt,
		DeletedAt:             ft.DeletedAt,
		Vendors:               Vendors,
	}
}
