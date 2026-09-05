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

type ClientDestinationService struct {
	Repo *repository.ClientDestinationRepository
}

func NewClientDestinationService(r *repository.ClientDestinationRepository) *ClientDestinationService {
	return &ClientDestinationService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimClientDestination(req *model.ClientDestinationRequest) {
	validation.TrimStrings(
		&req.ClientDestinationName,
		req.ClientDestinationMapUrl,
	)
}

// ==================================================
// Create
// ==================================================
func (s *ClientDestinationService) Create(ctx context.Context, loginID int, req *model.ClientDestinationRequest) (*model.ClientDestinationResponse, error, map[string]string) {
	trimClientDestination(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.ClientDestination{
		ClientID:                   req.ClientID,
		ClientDestinationName:      req.ClientDestinationName,
		CityID:                     req.CityID,
		ClientDestinationAddress:   req.ClientDestinationAddress,
		ClientDestinationLatitude:  req.ClientDestinationLatitude,
		ClientDestinationLongitude: req.ClientDestinationLongitude,
		ClientDestinationMapUrl:    req.ClientDestinationMapUrl,
		OperatingHours:             req.OperatingHours,
		IsActive:                   isActive,
		CreatedBy:                  loginID,
		UpdatedBy:                  loginID,
		CreatedAt:                  now,
		UpdatedAt:                  now,
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
func (s *ClientDestinationService) Update(ctx context.Context, loginID, id int, req *model.ClientDestinationRequest) (*model.ClientDestinationResponse, error, map[string]string) {
	trimClientDestination(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.ClientDestination{
		ClientID:                   req.ClientID,
		ClientDestinationName:      req.ClientDestinationName,
		CityID:                     req.CityID,
		ClientDestinationAddress:   req.ClientDestinationAddress,
		ClientDestinationLatitude:  req.ClientDestinationLatitude,
		ClientDestinationLongitude: req.ClientDestinationLongitude,
		ClientDestinationMapUrl:    req.ClientDestinationMapUrl,
		OperatingHours:             req.OperatingHours,
		UpdatedBy:                  loginID,
		UpdatedAt:                  time.Now(),
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
func (s *ClientDestinationService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *ClientDestinationService) GetByID(ctx context.Context, id int) (*model.ClientDestinationResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toClientDestinationResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *ClientDestinationService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.ClientDestinationResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.ClientDestinationResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toClientDestinationResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toClientDestinationResp(m *model.ClientDestination) *model.ClientDestinationResponse {
	if m == nil {
		return nil
	}

	var clientResp *model.ClientResponse
	var cityResp *model.CityResponse
	if m.Client != nil {
		clientResp = toClientResp(m.Client)
	}
	if m.City != nil {
		cityResp = toCityResp(m.City)
	}
	resp := &model.ClientDestinationResponse{
		ClientDestinationID:        m.ClientDestinationID,
		ClientID:                   m.ClientID,
		ClientDestinationName:      m.ClientDestinationName,
		CityID:                     m.CityID,
		ClientDestinationAddress:   m.ClientDestinationAddress,
		ClientDestinationLatitude:  m.ClientDestinationLatitude,
		ClientDestinationLongitude: m.ClientDestinationLongitude,
		ClientDestinationMapUrl:    m.ClientDestinationMapUrl,
		OperatingHours:             m.OperatingHours,
		IsActive:                   m.IsActive,
		CreatedAt:                  m.CreatedAt,
		UpdatedAt:                  m.UpdatedAt,
		DeletedAt:                  m.DeletedAt,
		Client:                     clientResp,
		City:                       cityResp,
	}

	return resp
}
