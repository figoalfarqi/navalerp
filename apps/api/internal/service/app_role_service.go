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

type AppRoleService struct {
	Repo *repository.AppRoleRepository
}

func NewAppRoleService(r *repository.AppRoleRepository) *AppRoleService {
	return &AppRoleService{Repo: r}
}

// ==================================================
// Utility (Trim String)
// ==================================================
func trimAppRole(req *model.AppRoleRequest) {
	validation.TrimStrings(&req.AppRoleName)
	if req.AppRoleDescription != nil {
		validation.TrimStrings(req.AppRoleDescription)
	}
}

// ==================================================
// Create
// ==================================================
func (s *AppRoleService) Create(ctx context.Context, loginID int, req *model.AppRoleRequest) (*model.AppRoleResponse, error, map[string]string) {
	trimAppRole(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	item := &model.AppRole{
		AppRoleTypeID:      req.AppRoleTypeID,
		AppRoleName:        req.AppRoleName,
		AppRoleDescription: req.AppRoleDescription,
		IsActive:           isActive,
		CreatedBy:          loginID,
		UpdatedBy:          loginID,
		CreatedAt:          now,
		UpdatedAt:          now,
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
func (s *AppRoleService) Update(ctx context.Context, loginID, id int, req *model.AppRoleRequest) (*model.AppRoleResponse, error, map[string]string) {
	trimAppRole(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	item := &model.AppRole{
		AppRoleTypeID:      req.AppRoleTypeID,
		AppRoleName:        req.AppRoleName,
		AppRoleDescription: req.AppRoleDescription,
		UpdatedBy:          loginID,
		UpdatedAt:          time.Now(),
	}

	if req.IsActive != nil {
		item.IsActive = *req.IsActive
	} else {
		item.IsActive = -1 // sentinel: repo harus abaikan update kolom is_active
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
func (s *AppRoleService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *AppRoleService) GetByID(ctx context.Context, id int) (*model.AppRoleResponse, error) {
	item, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toAppRoleResp(item), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *AppRoleService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.AppRoleResponse, error) {

	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.AppRoleResponse, 0, len(items))
	for _, d := range items {
		resp = append(resp, *toAppRoleResp(&d))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toAppRoleResp(m *model.AppRole) *model.AppRoleResponse {
	if m == nil {
		return nil
	}

	return &model.AppRoleResponse{
		AppRoleID:          m.AppRoleID,
		AppRoleTypeID:      m.AppRoleTypeID,
		AppRoleName:        m.AppRoleName,
		AppRoleDescription: m.AppRoleDescription,
		IsActive:           m.IsActive,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		DeletedAt:          m.DeletedAt,
	}
}
