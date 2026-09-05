package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type TheaterService struct {
	Repo *repository.TheaterRepository
}

func NewTheaterService(repo *repository.TheaterRepository) *TheaterService {
	return &TheaterService{Repo: repo}
}

func (s *TheaterService) GetByID(ctx context.Context, id string) (*model.Theater, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *TheaterService) List(ctx context.Context, opts model.ListOptions) ([]model.Theater, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *TheaterService) Create(ctx context.Context, loginID string, req *model.TheaterRequest) (*model.Theater, error, map[string]string) {
	m := &model.Theater{}
	if req.TheaterCode != nil { m.TheaterCode = *req.TheaterCode }
	if req.TheaterName != nil { m.TheaterName = *req.TheaterName }
	if req.ResponsibleCommandUnitId != nil { m.ResponsibleCommandUnitId = *req.ResponsibleCommandUnitId }
	m.ThreatLevel = req.ThreatLevel
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TheaterService) Update(ctx context.Context, loginID string, id string, req *model.TheaterRequest) (*model.Theater, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.TheaterCode != nil { m.TheaterCode = *req.TheaterCode }
	if req.TheaterName != nil { m.TheaterName = *req.TheaterName }
	if req.ResponsibleCommandUnitId != nil { m.ResponsibleCommandUnitId = *req.ResponsibleCommandUnitId }
	if req.ThreatLevel != nil { m.ThreatLevel = req.ThreatLevel }
	if req.Description != nil { m.Description = req.Description }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TheaterService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
