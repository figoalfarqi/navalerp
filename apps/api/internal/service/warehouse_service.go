package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type WarehouseService struct {
	Repo *repository.WarehouseRepository
}

func NewWarehouseService(repo *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{Repo: repo}
}

func (s *WarehouseService) GetByID(ctx context.Context, id string) (*model.Warehouse, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *WarehouseService) List(ctx context.Context, opts model.ListOptions) ([]model.Warehouse, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *WarehouseService) Create(ctx context.Context, loginID string, req *model.WarehouseRequest) (*model.Warehouse, error, map[string]string) {
	m := &model.Warehouse{}
	if req.UnitId != nil { m.UnitId = *req.UnitId }
	if req.WarehouseCode != nil { m.WarehouseCode = *req.WarehouseCode }
	if req.WarehouseName != nil { m.WarehouseName = *req.WarehouseName }
	if req.WarehouseType != nil { m.WarehouseType = *req.WarehouseType }
	m.CapacityM3 = req.CapacityM3
	m.ManagerUserId = req.ManagerUserId
	m.LocationAddress = req.LocationAddress
	m.IsActive = req.IsActive
	m.Locations = req.Locations
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *WarehouseService) Update(ctx context.Context, loginID string, id string, req *model.WarehouseRequest) (*model.Warehouse, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.UnitId != nil { m.UnitId = *req.UnitId }
	if req.WarehouseCode != nil { m.WarehouseCode = *req.WarehouseCode }
	if req.WarehouseName != nil { m.WarehouseName = *req.WarehouseName }
	if req.WarehouseType != nil { m.WarehouseType = *req.WarehouseType }
	if req.CapacityM3 != nil { m.CapacityM3 = req.CapacityM3 }
	if req.ManagerUserId != nil { m.ManagerUserId = req.ManagerUserId }
	if req.LocationAddress != nil { m.LocationAddress = req.LocationAddress }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if req.Locations != nil { m.Locations = req.Locations }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *WarehouseService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
