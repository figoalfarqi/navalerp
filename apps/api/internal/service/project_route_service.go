package service

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectRouteService struct {
	Repo *repository.ProjectRouteRepository
}

func NewProjectRouteService(repo *repository.ProjectRouteRepository) *ProjectRouteService {
	return &ProjectRouteService{Repo: repo}
}
func (s *ProjectRouteService) Create(ctx context.Context, userID int, req *model.ProjectRouteRequest) (*model.ProjectRoute, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectRouteService) Update(ctx context.Context, userID, id int, req *model.ProjectRouteRequest) (*model.ProjectRoute, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectRouteService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectRouteService) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectRoute, error) {
	return s.Repo.GetByID(ctx, id, checkerID)
}
func (s *ProjectRouteService) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectRoute, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
