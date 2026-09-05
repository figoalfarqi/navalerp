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

type CityService struct {
	Repo *repository.CityRepository
}

func NewCityService(r *repository.CityRepository) *CityService {
	return &CityService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimCity(req *model.CityRequest) {
	validation.TrimStrings(&req.CityName)
}

// ==================================================
// Create City
// ==================================================
func (s *CityService) Create(ctx context.Context, loginID int, req *model.CityRequest) (*model.CityResponse, error, map[string]string) {
	trimCity(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	// is_active default = 1
	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	city := &model.City{
		ProvinceID: req.ProvinceID,
		CityName:   req.CityName,
		IsActive:   isActive,
		CreatedBy:  loginID,
		UpdatedBy:  loginID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	id, err := s.Repo.Create(ctx, city)
	if err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Update City
// ==================================================
func (s *CityService) Update(ctx context.Context, loginID, id int, req *model.CityRequest) (*model.CityResponse, error, map[string]string) {
	trimCity(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	city := &model.City{
		ProvinceID: req.ProvinceID,
		CityName:   req.CityName,
		UpdatedBy:  loginID,
		UpdatedAt:  time.Now(),
	}

	// Jika IsActive tidak diisi: jangan update is_active
	if req.IsActive != nil {
		city.IsActive = *req.IsActive
	} else {
		city.IsActive = -1 // signal
	}

	if err := s.Repo.Update(ctx, id, city); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ==================================================
// Delete (Soft Delete)
// ==================================================
func (s *CityService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *CityService) GetByID(ctx context.Context, id int) (*model.CityResponse, error) {
	city, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toCityResp(city), nil
}

// ==================================================
// List with cursor + filters
// ==================================================
func (s *CityService) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string) ([]model.CityResponse, error) {

	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.CityResponse, 0, len(items))
	for _, c := range items {
		resp = append(resp, *toCityResp(&c))
	}

	return resp, nil
}

// ==================================================
// Mapping ke Response DTO
// ==================================================
func toCityResp(c *model.City) *model.CityResponse {
	if c == nil {
		return nil
	}

	var provinceResp *model.ProvinceResponse
	if c.Province != nil {
		provinceResp = toProvinceResp(c.Province)
	}

	return &model.CityResponse{
		CityID:     c.CityID,
		ProvinceID: c.ProvinceID,
		CityName:   c.CityName,
		IsActive:   c.IsActive,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		DeletedAt:  c.DeletedAt,
		Province:   provinceResp,
	}
}
