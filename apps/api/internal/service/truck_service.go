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

type TruckService struct {
	Repo *repository.TruckRepository
}

func NewTruckService(r *repository.TruckRepository) *TruckService {
	return &TruckService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimTruck(req *model.TruckRequest) {
	validation.TrimStrings(
		&req.LicensePlate,
	)
}

// ==================================================
// Create
// ==================================================
func (s *TruckService) Create(ctx context.Context, loginID int, req *model.TruckRequest) (*model.TruckResponse, error, map[string]string) {
	trimTruck(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.Truck{
		TruckTypeID:       req.TruckTypeID,
		TruckMerkID:       req.TruckMerkID,
		DriverID:          req.DriverID,
		VendorID:          req.VendorID,
		LicensePlate:      req.LicensePlate,
		OwnershipStatusID: req.OwnershipStatusID,
		ProductionYear:    req.ProductionYear,
		NumberOfTires:     req.NumberOfTires,
		IsActive:          isActive,
		CreatedBy:         loginID,
		UpdatedBy:         loginID,
		CreatedAt:         now,
		UpdatedAt:         now,
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
func (s *TruckService) Update(ctx context.Context, loginID, id int, req *model.TruckRequest) (*model.TruckResponse, error, map[string]string) {
	trimTruck(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.Truck{
		TruckTypeID:       req.TruckTypeID,
		TruckMerkID:       req.TruckMerkID,
		DriverID:          req.DriverID,
		VendorID:          req.VendorID,
		LicensePlate:      req.LicensePlate,
		OwnershipStatusID: req.OwnershipStatusID,
		ProductionYear:    req.ProductionYear,
		NumberOfTires:     req.NumberOfTires,
		UpdatedBy:         loginID,
		UpdatedAt:         time.Now(),
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
func (s *TruckService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *TruckService) GetByID(ctx context.Context, id int) (*model.TruckResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTruckResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *TruckService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.TruckResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.TruckResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toTruckResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toTruckResp(m *model.Truck) *model.TruckResponse {
	if m == nil {
		return nil
	}

	var truckTypeResp *model.TruckTypeResponse
	if m.TruckType != nil {
		truckTypeResp = toTruckTypeResp(m.TruckType)
	}
	var truckMerkResp *model.TruckMerkResponse
	if m.TruckMerk != nil {
		truckMerkResp = toTruckMerkResp(m.TruckMerk)
	}
	var driverResp *model.AppUserResponse
	if m.Driver != nil {
		driverResp = toAppUserResp(m.Driver)
	}
	var vendorResp *model.VendorResponse
	if m.Vendor != nil {
		vendorResp = toVendorResp(m.Vendor)
	}
	return &model.TruckResponse{
		TruckID:           m.TruckID,
		TruckTypeID:       m.TruckTypeID,
		TruckMerkID:       m.TruckMerkID,
		DriverID:          m.DriverID,
		VendorID:          m.VendorID,
		LicensePlate:      m.LicensePlate,
		OwnershipStatusID: m.OwnershipStatusID,
		ProductionYear:    m.ProductionYear,
		NumberOfTires:     m.NumberOfTires,
		IsActive:          m.IsActive,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         m.DeletedAt,
		TruckType:         truckTypeResp,
		TruckMerk:         truckMerkResp,
		Driver:            driverResp,
		Vendor:            vendorResp,
	}
}
