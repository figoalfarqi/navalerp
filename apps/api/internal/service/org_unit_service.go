package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type OrgUnitService struct {
	Repo *repository.OrgUnitRepository
}

func NewOrgUnitService(repo *repository.OrgUnitRepository) *OrgUnitService {
	return &OrgUnitService{Repo: repo}
}

func (s *OrgUnitService) GetByID(ctx context.Context, id string) (*model.OrgUnit, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *OrgUnitService) List(ctx context.Context, opts model.ListOptions) ([]model.OrgUnit, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *OrgUnitService) Create(ctx context.Context, loginID string, req *model.OrgUnitRequest) (*model.OrgUnit, error, map[string]string) {
	m := &model.OrgUnit{}
	m.ParentUnitId = req.ParentUnitId
	if req.UnitCode != nil { m.UnitCode = *req.UnitCode }
	if req.UnitName != nil { m.UnitName = *req.UnitName }
	if req.UnitType != nil { m.UnitType = *req.UnitType }
	m.Description = req.Description
	m.CommandLevel = req.CommandLevel
	m.Latitude = req.Latitude
	m.Longitude = req.Longitude
	m.Address = req.Address
	m.Phone = req.Phone
	m.IsActive = req.IsActive
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *OrgUnitService) Update(ctx context.Context, loginID string, id string, req *model.OrgUnitRequest) (*model.OrgUnit, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ParentUnitId != nil { m.ParentUnitId = req.ParentUnitId }
	if req.UnitCode != nil { m.UnitCode = *req.UnitCode }
	if req.UnitName != nil { m.UnitName = *req.UnitName }
	if req.UnitType != nil { m.UnitType = *req.UnitType }
	if req.Description != nil { m.Description = req.Description }
	if req.CommandLevel != nil { m.CommandLevel = req.CommandLevel }
	if req.Latitude != nil { m.Latitude = req.Latitude }
	if req.Longitude != nil { m.Longitude = req.Longitude }
	if req.Address != nil { m.Address = req.Address }
	if req.Phone != nil { m.Phone = req.Phone }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *OrgUnitService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
