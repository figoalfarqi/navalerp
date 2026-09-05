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

type CargoTypeService struct {
	Repo *repository.CargoTypeRepository
}

func NewCargoTypeService(r *repository.CargoTypeRepository) *CargoTypeService {
	return &CargoTypeService{Repo: r}
}

// ==================================================
// Utility (Trim String)
// ==================================================
func trimCargoType(req *model.CargoTypeRequest) {
	validation.TrimStrings(&req.CargoTypeName, req.CargoTypeDescription)
}

// ==================================================
// Create
// ==================================================
func (s *CargoTypeService) Create(ctx context.Context, loginID int, req *model.CargoTypeRequest) (*model.CargoTypeResponse, error, map[string]string) {
	trimCargoType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	item := &model.CargoType{
		CargoTypeName:        req.CargoTypeName,
		CargoTypeDescription: req.CargoTypeDescription,
		CargoTypeGrade:       req.CargoTypeGrade,
		IsActive:             isActive,
		CreatedBy:            loginID,
		UpdatedBy:            loginID,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	id, err := s.Repo.Create(ctx, item)
	if err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Update
// ==================================================
func (s *CargoTypeService) Update(ctx context.Context, loginID, id int, req *model.CargoTypeRequest) (*model.CargoTypeResponse, error, map[string]string) {
	trimCargoType(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	item := &model.CargoType{
		CargoTypeName:        req.CargoTypeName,
		CargoTypeDescription: req.CargoTypeDescription,
		CargoTypeGrade:       req.CargoTypeGrade,
		UpdatedBy:            loginID,
		UpdatedAt:            time.Now(),
	}

	// → Sama dengan ClientService:
	// Jika req.IsActive tidak dikirim → jangan update kolom is_active
	if req.IsActive != nil {
		item.IsActive = *req.IsActive
	} else {
		item.IsActive = -1 // sentinel → repo akan tahu: jangan update kolom ini
	}

	if err := s.Repo.Update(ctx, id, item); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Delete (Soft Delete)
// ==================================================
func (s *CargoTypeService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *CargoTypeService) GetByID(ctx context.Context, id int) (*model.CargoTypeResponse, error) {
	item, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toCargoTypeResp(item), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *CargoTypeService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.CargoTypeResponse, error) {
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.CargoTypeResponse, 0, len(items))
	for _, d := range items {
		resp = append(resp, *toCargoTypeResp(&d))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toCargoTypeResp(m *model.CargoType) *model.CargoTypeResponse {
	if m == nil {
		return nil
	}

	return &model.CargoTypeResponse{
		CargoTypeID:          m.CargoTypeID,
		CargoTypeName:        m.CargoTypeName,
		CargoTypeDescription: m.CargoTypeDescription,
		CargoTypeGrade:       m.CargoTypeGrade,
		IsActive:             m.IsActive,
		CreatedAt:            m.CreatedAt,
		UpdatedAt:            m.UpdatedAt,
		DeletedAt:            m.DeletedAt,
	}
}
