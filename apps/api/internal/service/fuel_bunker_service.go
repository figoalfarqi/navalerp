package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type FuelBunkerService struct {
	Repo *repository.FuelBunkerRepository
}

func NewFuelBunkerService(repo *repository.FuelBunkerRepository) *FuelBunkerService {
	return &FuelBunkerService{Repo: repo}
}

func (s *FuelBunkerService) GetByID(ctx context.Context, id string) (*model.FuelBunker, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *FuelBunkerService) List(ctx context.Context, opts model.ListOptions) ([]model.FuelBunker, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *FuelBunkerService) Create(ctx context.Context, loginID string, req *model.FuelBunkerRequest) (*model.FuelBunker, error, map[string]string) {
	m := &model.FuelBunker{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	m.FacilityId = req.FacilityId
	if req.FuelType != nil { m.FuelType = *req.FuelType }
	if req.QuantityLiters != nil { m.QuantityLiters = *req.QuantityLiters }
	m.Density15c = req.Density15c
	m.FlowRateLph = req.FlowRateLph
	if req.BunkeringStartTime != nil { m.BunkeringStartTime = *req.BunkeringStartTime }
	if req.BunkeringEndTime != nil { m.BunkeringEndTime = *req.BunkeringEndTime }
	if req.ReceiptVoucherNo != nil { m.ReceiptVoucherNo = *req.ReceiptVoucherNo }
	if req.AuthorisedByUserId != nil { m.AuthorisedByUserId = *req.AuthorisedByUserId }
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *FuelBunkerService) Update(ctx context.Context, loginID string, id string, req *model.FuelBunkerRequest) (*model.FuelBunker, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.FacilityId != nil { m.FacilityId = req.FacilityId }
	if req.FuelType != nil { m.FuelType = *req.FuelType }
	if req.QuantityLiters != nil { m.QuantityLiters = *req.QuantityLiters }
	if req.Density15c != nil { m.Density15c = req.Density15c }
	if req.FlowRateLph != nil { m.FlowRateLph = req.FlowRateLph }
	if req.BunkeringStartTime != nil { m.BunkeringStartTime = *req.BunkeringStartTime }
	if req.BunkeringEndTime != nil { m.BunkeringEndTime = *req.BunkeringEndTime }
	if req.ReceiptVoucherNo != nil { m.ReceiptVoucherNo = *req.ReceiptVoucherNo }
	if req.AuthorisedByUserId != nil { m.AuthorisedByUserId = *req.AuthorisedByUserId }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *FuelBunkerService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
