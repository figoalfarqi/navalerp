package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ShipClassService struct {
	Repo *repository.ShipClassRepository
}

func NewShipClassService(repo *repository.ShipClassRepository) *ShipClassService {
	return &ShipClassService{Repo: repo}
}

func (s *ShipClassService) GetByID(ctx context.Context, id string) (*model.ShipClass, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ShipClassService) List(ctx context.Context, opts model.ListOptions) ([]model.ShipClass, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ShipClassService) Create(ctx context.Context, loginID string, req *model.ShipClassRequest) (*model.ShipClass, error, map[string]string) {
	m := &model.ShipClass{}
	if req.ClassCode != nil { m.ClassCode = *req.ClassCode }
	if req.ClassName != nil { m.ClassName = *req.ClassName }
	if req.Category != nil { m.Category = *req.Category }
	if req.Specifications != nil { m.Specifications = req.Specifications }
	m.Builder = req.Builder
	m.TotalBuilt = req.TotalBuilt
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipClassService) Update(ctx context.Context, loginID string, id string, req *model.ShipClassRequest) (*model.ShipClass, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ClassCode != nil { m.ClassCode = *req.ClassCode }
	if req.ClassName != nil { m.ClassName = *req.ClassName }
	if req.Category != nil { m.Category = *req.Category }
	if req.Specifications != nil { m.Specifications = req.Specifications }
	if req.Builder != nil { m.Builder = req.Builder }
	if req.TotalBuilt != nil { m.TotalBuilt = req.TotalBuilt }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipClassService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
