package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type RouteService struct {
	Repo *repository.RouteRepository
}

func NewRouteService(repo *repository.RouteRepository) *RouteService {
	return &RouteService{Repo: repo}
}

func (s *RouteService) GetByID(ctx context.Context, id string) (*model.Route, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *RouteService) List(ctx context.Context, opts model.ListOptions) ([]model.Route, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *RouteService) Create(ctx context.Context, loginID string, req *model.RouteRequest) (*model.Route, error, map[string]string) {
	m := &model.Route{}
	if req.RouteCode != nil { m.RouteCode = *req.RouteCode }
	if req.RouteName != nil { m.RouteName = *req.RouteName }
	if req.OriginFacilityId != nil { m.OriginFacilityId = *req.OriginFacilityId }
	if req.DestinationFacilityId != nil { m.DestinationFacilityId = *req.DestinationFacilityId }
	if req.DistanceNauticalMiles != nil { m.DistanceNauticalMiles = *req.DistanceNauticalMiles }
	if req.EstimatedTransitHours != nil { m.EstimatedTransitHours = *req.EstimatedTransitHours }
	m.RiskLevel = req.RiskLevel
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *RouteService) Update(ctx context.Context, loginID string, id string, req *model.RouteRequest) (*model.Route, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.RouteCode != nil { m.RouteCode = *req.RouteCode }
	if req.RouteName != nil { m.RouteName = *req.RouteName }
	if req.OriginFacilityId != nil { m.OriginFacilityId = *req.OriginFacilityId }
	if req.DestinationFacilityId != nil { m.DestinationFacilityId = *req.DestinationFacilityId }
	if req.DistanceNauticalMiles != nil { m.DistanceNauticalMiles = *req.DistanceNauticalMiles }
	if req.EstimatedTransitHours != nil { m.EstimatedTransitHours = *req.EstimatedTransitHours }
	if req.RiskLevel != nil { m.RiskLevel = req.RiskLevel }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *RouteService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
