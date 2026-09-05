package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type TransportUnitService struct {
	Repo *repository.TransportUnitRepository
}

func NewTransportUnitService(repo *repository.TransportUnitRepository) *TransportUnitService {
	return &TransportUnitService{Repo: repo}
}

func (s *TransportUnitService) GetByID(ctx context.Context, id string) (*model.TransportUnit, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *TransportUnitService) List(ctx context.Context, opts model.ListOptions) ([]model.TransportUnit, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *TransportUnitService) Create(ctx context.Context, loginID string, req *model.TransportUnitRequest) (*model.TransportUnit, error, map[string]string) {
	m := &model.TransportUnit{}
	if req.UnitCode != nil { m.UnitCode = *req.UnitCode }
	if req.UnitName != nil { m.UnitName = *req.UnitName }
	if req.TransportType != nil { m.TransportType = *req.TransportType }
	if req.CargoCapacityTons != nil { m.CargoCapacityTons = *req.CargoCapacityTons }
	m.FuelCapacityLiters = req.FuelCapacityLiters
	if req.OperatingUnitId != nil { m.OperatingUnitId = *req.OperatingUnitId }
	m.Status = req.Status
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TransportUnitService) Update(ctx context.Context, loginID string, id string, req *model.TransportUnitRequest) (*model.TransportUnit, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.UnitCode != nil { m.UnitCode = *req.UnitCode }
	if req.UnitName != nil { m.UnitName = *req.UnitName }
	if req.TransportType != nil { m.TransportType = *req.TransportType }
	if req.CargoCapacityTons != nil { m.CargoCapacityTons = *req.CargoCapacityTons }
	if req.FuelCapacityLiters != nil { m.FuelCapacityLiters = req.FuelCapacityLiters }
	if req.OperatingUnitId != nil { m.OperatingUnitId = *req.OperatingUnitId }
	if req.Status != nil { m.Status = req.Status }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TransportUnitService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
