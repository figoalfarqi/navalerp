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

type ProvinceService struct {
	Repo *repository.ProvinceRepository
}

func NewProvinceService(r *repository.ProvinceRepository) *ProvinceService {
	return &ProvinceService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimProvince(req *model.ProvinceRequest) {
	validation.TrimStrings(&req.ProvinceName, req.ProvinceRealName)
}

// ==================================================
// Create Province
// ==================================================
func (s *ProvinceService) Create(
	ctx context.Context,
	loginID int,
	req *model.ProvinceRequest,
) (*model.ProvinceResponse, error, map[string]string) {

	trimProvince(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	p := &model.Province{
		ProvinceName:     req.ProvinceName,
		ProvinceRealName: req.ProvinceRealName,
		CreatedBy:        loginID,
		UpdatedBy:        loginID,
		CreatedAt:        now,
		UpdatedAt:        now,
		IsActive:         1, // default aktif
	}

	id, err := s.Repo.Create(ctx, p)
	if err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Update Province
// ==================================================
func (s *ProvinceService) Update(
	ctx context.Context,
	loginID, id int,
	req *model.ProvinceRequest,
) (*model.ProvinceResponse, error, map[string]string) {

	trimProvince(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	p := &model.Province{
		ProvinceName:     req.ProvinceName,
		ProvinceRealName: req.ProvinceRealName,
		UpdatedBy:        loginID,
		UpdatedAt:        time.Now(),
	}

	if req.IsActive != nil {
		p.IsActive = *req.IsActive
	} else {
		p.IsActive = -1 // jangan update kolom is_active
	}

	if err := s.Repo.Update(ctx, id, p); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Delete (Soft Delete)
// ==================================================
func (s *ProvinceService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *ProvinceService) GetByID(ctx context.Context, id int) (*model.ProvinceResponse, error) {
	p, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toProvinceResp(p), nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (s *ProvinceService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.ProvinceResponse, error) {

	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.ProvinceResponse, 0, len(items))
	for _, p := range items {
		resp = append(resp, *toProvinceResp(&p))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toProvinceResp(p *model.Province) *model.ProvinceResponse {
	if p == nil {
		return nil
	}

	Cities := make([]model.CityResponse, 0)
	if p.Cities != nil {
		for _, cty := range p.Cities {
			Cities = append(Cities, *toCityResp(&cty))
		}
	}
	return &model.ProvinceResponse{
		ProvinceID:       p.ProvinceID,
		ProvinceName:     p.ProvinceName,
		ProvinceRealName: p.ProvinceRealName,
		IsActive:         p.IsActive,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
		DeletedAt:        p.DeletedAt,
		Cities:           Cities,
	}
}
