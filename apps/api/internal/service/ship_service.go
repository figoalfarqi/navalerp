package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ShipService struct {
	Repo *repository.ShipRepository
}

func NewShipService(repo *repository.ShipRepository) *ShipService {
	return &ShipService{Repo: repo}
}

func (s *ShipService) GetByID(ctx context.Context, id string) (*model.Ship, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ShipService) List(ctx context.Context, opts model.ListOptions) ([]model.Ship, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ShipService) Create(ctx context.Context, loginID string, req *model.ShipRequest) (*model.Ship, error, map[string]string) {
	m := &model.Ship{}
	if req.ClassId != nil { m.ClassId = *req.ClassId }
	if req.AssignedUnitId != nil { m.AssignedUnitId = *req.AssignedUnitId }
	if req.HullNumber != nil { m.HullNumber = *req.HullNumber }
	if req.ShipName != nil { m.ShipName = *req.ShipName }
	m.CallSign = req.CallSign
	m.CommissionDate = req.CommissionDate
	m.HomePort = req.HomePort
	m.LengthM = req.LengthM
	m.BeamM = req.BeamM
	m.DraftM = req.DraftM
	m.DisplacementTons = req.DisplacementTons
	m.MaxSpeedKnots = req.MaxSpeedKnots
	m.CruiseRangeNm = req.CruiseRangeNm
	m.CrewCapacity = req.CrewCapacity
	m.FuelCapacityLiters = req.FuelCapacityLiters
	m.FreshWaterCapacityLiters = req.FreshWaterCapacityLiters
	m.Status = req.Status
	m.CurrentReadinessStatus = req.CurrentReadinessStatus
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipService) Update(ctx context.Context, loginID string, id string, req *model.ShipRequest) (*model.Ship, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ClassId != nil { m.ClassId = *req.ClassId }
	if req.AssignedUnitId != nil { m.AssignedUnitId = *req.AssignedUnitId }
	if req.HullNumber != nil { m.HullNumber = *req.HullNumber }
	if req.ShipName != nil { m.ShipName = *req.ShipName }
	if req.CallSign != nil { m.CallSign = req.CallSign }
	if req.CommissionDate != nil { m.CommissionDate = req.CommissionDate }
	if req.HomePort != nil { m.HomePort = req.HomePort }
	if req.LengthM != nil { m.LengthM = req.LengthM }
	if req.BeamM != nil { m.BeamM = req.BeamM }
	if req.DraftM != nil { m.DraftM = req.DraftM }
	if req.DisplacementTons != nil { m.DisplacementTons = req.DisplacementTons }
	if req.MaxSpeedKnots != nil { m.MaxSpeedKnots = req.MaxSpeedKnots }
	if req.CruiseRangeNm != nil { m.CruiseRangeNm = req.CruiseRangeNm }
	if req.CrewCapacity != nil { m.CrewCapacity = req.CrewCapacity }
	if req.FuelCapacityLiters != nil { m.FuelCapacityLiters = req.FuelCapacityLiters }
	if req.FreshWaterCapacityLiters != nil { m.FreshWaterCapacityLiters = req.FreshWaterCapacityLiters }
	if req.Status != nil { m.Status = req.Status }
	if req.CurrentReadinessStatus != nil { m.CurrentReadinessStatus = req.CurrentReadinessStatus }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
