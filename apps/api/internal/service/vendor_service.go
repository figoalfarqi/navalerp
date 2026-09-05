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

type VendorService struct {
	Repo *repository.VendorRepository
}

func NewVendorService(r *repository.VendorRepository) *VendorService {
	return &VendorService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimVendor(req *model.VendorRequest) {
	validation.TrimStrings(
		&req.VendorName,
	)
}

// ==================================================
// Create
// ==================================================
func (s *VendorService) Create(ctx context.Context, loginID int, req *model.VendorRequest) (*model.VendorResponse, error, map[string]string) {
	trimVendor(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.Vendor{
		VendorTypeID:      req.VendorTypeID,
		BankMerkID:        req.BankMerkID,
		VendorName:        req.VendorName,
		VendorEmail:       req.VendorEmail,
		VendorPhone:       req.VendorPhone,
		VendorTin:         req.VendorTin,
		CityID:            req.CityID,
		VendorAddress:     req.VendorAddress,
		BankAccountNumber: req.BankAccountNumber,
		BankAccountName:   req.BankAccountName,
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
func (s *VendorService) Update(ctx context.Context, loginID, id int, req *model.VendorRequest) (*model.VendorResponse, error, map[string]string) {
	trimVendor(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.Vendor{
		VendorTypeID:      req.VendorTypeID,
		BankMerkID:        req.BankMerkID,
		VendorName:        req.VendorName,
		VendorEmail:       req.VendorEmail,
		VendorPhone:       req.VendorPhone,
		VendorTin:         req.VendorTin,
		CityID:            req.CityID,
		VendorAddress:     req.VendorAddress,
		BankAccountNumber: req.BankAccountNumber,
		BankAccountName:   req.BankAccountName,
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
func (s *VendorService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *VendorService) GetByID(ctx context.Context, id int) (*model.VendorResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toVendorResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *VendorService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.VendorResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.VendorResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toVendorResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toVendorResp(m *model.Vendor) *model.VendorResponse {
	if m == nil {
		return nil
	}

	var vendorTypeResp *model.VendorTypeResponse
	if m.VendorType != nil {
		vendorTypeResp = toVendorTypeResp(m.VendorType)
	}
	var bankMerkResp *model.BankMerkResponse
	if m.BankMerk != nil {
		bankMerkResp = toBankMerkResp(m.BankMerk)
	}
	var cityResp *model.CityResponse
	if m.City != nil {
		cityResp = toCityResp(m.City)
	}
	Trucks := make([]model.TruckResponse, 0)
	if m.Trucks != nil {
		for _, truck := range m.Trucks {
			Trucks = append(Trucks, *toTruckResp(&truck))
		}
	}
	resp := &model.VendorResponse{
		VendorID:          m.VendorID,
		VendorTypeID:      m.VendorTypeID,
		BankMerkID:        m.BankMerkID,
		VendorName:        m.VendorName,
		VendorEmail:       m.VendorEmail,
		VendorPhone:       m.VendorPhone,
		VendorTin:         m.VendorTin,
		CityID:            m.CityID,
		VendorAddress:     m.VendorAddress,
		BankAccountNumber: m.BankAccountNumber,
		BankAccountName:   m.BankAccountName,
		IsActive:          m.IsActive,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         m.DeletedAt,
		VendorType:        vendorTypeResp,
		BankMerk:          bankMerkResp,
		City:              cityResp,
		Trucks:            Trucks,
	}
	return resp
}
