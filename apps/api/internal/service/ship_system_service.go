package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ShipSystemService struct {
	Repo *repository.ShipSystemRepository
}

func NewShipSystemService(repo *repository.ShipSystemRepository) *ShipSystemService {
	return &ShipSystemService{Repo: repo}
}

func (s *ShipSystemService) GetByID(ctx context.Context, id string) (*model.ShipSystem, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ShipSystemService) List(ctx context.Context, opts model.ListOptions) ([]model.ShipSystem, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ShipSystemService) Create(ctx context.Context, loginID string, req *model.ShipSystemRequest) (*model.ShipSystem, error, map[string]string) {
	m := &model.ShipSystem{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	m.ParentSystemId = req.ParentSystemId
	if req.SystemCode != nil { m.SystemCode = *req.SystemCode }
	if req.SystemName != nil { m.SystemName = *req.SystemName }
	if req.SystemCategory != nil { m.SystemCategory = *req.SystemCategory }
	if req.SystemLevel != nil { m.SystemLevel = *req.SystemLevel }
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipSystemService) Update(ctx context.Context, loginID string, id string, req *model.ShipSystemRequest) (*model.ShipSystem, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.ParentSystemId != nil { m.ParentSystemId = req.ParentSystemId }
	if req.SystemCode != nil { m.SystemCode = *req.SystemCode }
	if req.SystemName != nil { m.SystemName = *req.SystemName }
	if req.SystemCategory != nil { m.SystemCategory = *req.SystemCategory }
	if req.SystemLevel != nil { m.SystemLevel = *req.SystemLevel }
	if req.Description != nil { m.Description = req.Description }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipSystemService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
