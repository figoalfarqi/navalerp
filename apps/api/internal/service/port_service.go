package service

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PortService struct{ Repo *repository.PortRepository }

func NewPortService(repo *repository.PortRepository) *PortService { return &PortService{Repo: repo} }

func (s *PortService) Create(ctx context.Context, userID int, req *model.PortRequest) (*model.Port, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}

func (s *PortService) Update(ctx context.Context, userID, id int, req *model.PortRequest) (*model.Port, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}

func (s *PortService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *PortService) GetByID(ctx context.Context, id int) (*model.Port, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *PortService) List(ctx context.Context, opts model.ListOptions) ([]model.Port, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
