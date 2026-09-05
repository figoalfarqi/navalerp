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

type ClientService struct {
	Repo *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{Repo: r}
}

// ==================================================
// Utility
// ==================================================
func trimClient(req *model.ClientRequest) {
	validation.TrimStrings(
		&req.ClientName,
		req.ClientEmail,
		req.ClientTin,
		req.ClientAddress,
	)
}

// ==================================================
// Create
// ==================================================
func (s *ClientService) Create(ctx context.Context, loginID int, req *model.ClientRequest) (*model.ClientResponse, error, map[string]string) {
	trimClient(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	now := time.Now()

	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	m := &model.Client{
		ClientName:          req.ClientName,
		ClientEmail:         req.ClientEmail,
		ClientTin:           req.ClientTin,
		NumberOfDayUntilDue: req.NumberOfDayUntilDue,
		CityID:              req.CityID,
		ClientAddress:       req.ClientAddress,
		OperatingHours:      req.OperatingHours,
		IsActive:            isActive,
		CreatedBy:           loginID,
		UpdatedBy:           loginID,
		CreatedAt:           now,
		UpdatedAt:           now,
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
func (s *ClientService) Update(ctx context.Context, loginID, id int, req *model.ClientRequest) (*model.ClientResponse, error, map[string]string) {
	trimClient(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	m := &model.Client{
		ClientName:          req.ClientName,
		ClientEmail:         req.ClientEmail,
		ClientTin:           req.ClientTin,
		NumberOfDayUntilDue: req.NumberOfDayUntilDue,
		CityID:              req.CityID,
		ClientAddress:       req.ClientAddress,
		OperatingHours:      req.OperatingHours,
		UpdatedBy:           loginID,
		UpdatedAt:           time.Now(),
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
func (s *ClientService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ==================================================
// Get By ID
// ==================================================
func (s *ClientService) GetByID(ctx context.Context, id int) (*model.ClientResponse, error) {
	m, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toClientResp(m), nil
}

// ==================================================
// List (cursor pagination + filters)
// ==================================================
func (s *ClientService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.ClientResponse, error) {
	// Normalisasi date filter seperti: created_at_after, created_at_before
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.ClientResponse, 0, len(items))
	for _, m := range items {
		resp = append(resp, *toClientResp(&m))
	}

	return resp, nil
}

// ==================================================
// Mapping ke response DTO
// ==================================================
func toClientResp(m *model.Client) *model.ClientResponse {
	if m == nil {
		return nil
	}

	var cityResp *model.CityResponse
	if m.City != nil {
		cityResp = toCityResp(m.City)
	}
	ClientDestinations := make([]model.ClientDestinationResponse, 0)
	if m.ClientDestinations != nil {
		for _, clndes := range m.ClientDestinations {
			ClientDestinations = append(ClientDestinations, *toClientDestinationResp(&clndes))
		}
	}

	AppUsers := make([]model.AppUserResponse, 0)
	if m.ClientPics != nil {
		for _, appuser := range m.ClientPics {
			AppUsers = append(AppUsers, *toAppUserResp(&appuser))
		}
	}

	resp := &model.ClientResponse{
		ClientID:            m.ClientID,
		ClientName:          m.ClientName,
		ClientEmail:         m.ClientEmail,
		ClientTin:           m.ClientTin,
		NumberOfDayUntilDue: m.NumberOfDayUntilDue,
		CityID:              m.CityID,
		ClientAddress:       m.ClientAddress,
		OperatingHours:      m.OperatingHours,
		IsActive:            m.IsActive,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
		DeletedAt:           m.DeletedAt,
		City:                cityResp,
		ClientDestinations:  ClientDestinations,
		ClientPics:          AppUsers,
	}

	return resp
}
