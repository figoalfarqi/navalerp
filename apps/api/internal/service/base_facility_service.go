package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type BaseFacilityService struct {
	Repo *repository.BaseFacilityRepository
}

func NewBaseFacilityService(repo *repository.BaseFacilityRepository) *BaseFacilityService {
	return &BaseFacilityService{Repo: repo}
}

func (s *BaseFacilityService) GetByID(ctx context.Context, id string) (*model.BaseFacility, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *BaseFacilityService) List(ctx context.Context, opts model.ListOptions) ([]model.BaseFacility, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *BaseFacilityService) Create(ctx context.Context, loginID string, req *model.BaseFacilityRequest) (*model.BaseFacility, error, map[string]string) {
	m := &model.BaseFacility{}
	if req.BaseUnitId != nil { m.BaseUnitId = *req.BaseUnitId }
	if req.FacilityCode != nil { m.FacilityCode = *req.FacilityCode }
	if req.FacilityName != nil { m.FacilityName = *req.FacilityName }
	if req.FacilityType != nil { m.FacilityType = *req.FacilityType }
	m.LengthMeters = req.LengthMeters
	m.DraftDepthMeters = req.DraftDepthMeters
	m.MaxDisplacementTonnage = req.MaxDisplacementTonnage
	m.HasShorePower = req.HasShorePower
	m.HasFreshWater = req.HasFreshWater
	m.HasFuelBunkerLine = req.HasFuelBunkerLine
	m.Status = req.Status
	m.Maintenances = req.Maintenances
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BaseFacilityService) Update(ctx context.Context, loginID string, id string, req *model.BaseFacilityRequest) (*model.BaseFacility, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.BaseUnitId != nil { m.BaseUnitId = *req.BaseUnitId }
	if req.FacilityCode != nil { m.FacilityCode = *req.FacilityCode }
	if req.FacilityName != nil { m.FacilityName = *req.FacilityName }
	if req.FacilityType != nil { m.FacilityType = *req.FacilityType }
	if req.LengthMeters != nil { m.LengthMeters = req.LengthMeters }
	if req.DraftDepthMeters != nil { m.DraftDepthMeters = req.DraftDepthMeters }
	if req.MaxDisplacementTonnage != nil { m.MaxDisplacementTonnage = req.MaxDisplacementTonnage }
	if req.HasShorePower != nil { m.HasShorePower = req.HasShorePower }
	if req.HasFreshWater != nil { m.HasFreshWater = req.HasFreshWater }
	if req.HasFuelBunkerLine != nil { m.HasFuelBunkerLine = req.HasFuelBunkerLine }
	if req.Status != nil { m.Status = req.Status }
	if req.Maintenances != nil { m.Maintenances = req.Maintenances }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BaseFacilityService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
