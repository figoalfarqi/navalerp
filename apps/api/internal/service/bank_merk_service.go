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

type BankMerkService struct {
	Repo *repository.BankMerkRepository
}

func NewBankMerkService(r *repository.BankMerkRepository) *BankMerkService {
	return &BankMerkService{Repo: r}
}

// ==================================================
// Utility (Trim String)
// ==================================================
func trimBankMerk(req *model.BankMerkRequest) {
	validation.TrimStrings(&req.BankMerkName)
	if req.BankMerkDescription != nil {
		validation.TrimStrings(req.BankMerkDescription)
	}
}

// ==================================================
// Create
// ==================================================
func (s *BankMerkService) Create(ctx context.Context, loginID int, req *model.BankMerkRequest) (*model.BankMerkResponse, error, map[string]string) {
	trimBankMerk(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	item := &model.BankMerk{
		BankMerkName:        req.BankMerkName,
		BankMerkDescription: req.BankMerkDescription,
		IsActive:            isActive,
		CreatedBy:           loginID,
		UpdatedBy:           loginID,
		CreatedAt:           now,
		UpdatedAt:           now,
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
func (s *BankMerkService) Update(ctx context.Context, loginID, id int, req *model.BankMerkRequest) (*model.BankMerkResponse, error, map[string]string) {
	trimBankMerk(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	item := &model.BankMerk{
		BankMerkName:        req.BankMerkName,
		BankMerkDescription: req.BankMerkDescription,
		UpdatedBy:           loginID,
		UpdatedAt:           time.Now(),
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
func (s *BankMerkService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *BankMerkService) GetByID(ctx context.Context, id int) (*model.BankMerkResponse, error) {
	item, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toBankMerkResp(item), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *BankMerkService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.BankMerkResponse, error) {
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.BankMerkResponse, 0, len(items))
	for _, d := range items {
		resp = append(resp, *toBankMerkResp(&d))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toBankMerkResp(m *model.BankMerk) *model.BankMerkResponse {
	if m == nil {
		return nil
	}

	return &model.BankMerkResponse{
		BankMerkID:          m.BankMerkID,
		BankMerkName:        m.BankMerkName,
		BankMerkDescription: m.BankMerkDescription,
		IsActive:            m.IsActive,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
		DeletedAt:           m.DeletedAt,
	}
}
