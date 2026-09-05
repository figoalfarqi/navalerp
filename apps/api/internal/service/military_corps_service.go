package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type MilitaryCorpsService struct {
	Repo *repository.MilitaryCorpsRepository
}

func NewMilitaryCorpsService(repo *repository.MilitaryCorpsRepository) *MilitaryCorpsService {
	return &MilitaryCorpsService{Repo: repo}
}

func (s *MilitaryCorpsService) GetByID(ctx context.Context, id string) (*model.MilitaryCorps, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *MilitaryCorpsService) List(ctx context.Context, opts model.ListOptions) ([]model.MilitaryCorps, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *MilitaryCorpsService) Create(ctx context.Context, loginID string, req *model.MilitaryCorpsRequest) (*model.MilitaryCorps, error, map[string]string) {
	m := &model.MilitaryCorps{}
	if req.CorpsCode != nil { m.CorpsCode = *req.CorpsCode }
	if req.CorpsName != nil { m.CorpsName = *req.CorpsName }
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MilitaryCorpsService) Update(ctx context.Context, loginID string, id string, req *model.MilitaryCorpsRequest) (*model.MilitaryCorps, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.CorpsCode != nil { m.CorpsCode = *req.CorpsCode }
	if req.CorpsName != nil { m.CorpsName = *req.CorpsName }
	if req.Description != nil { m.Description = req.Description }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MilitaryCorpsService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
