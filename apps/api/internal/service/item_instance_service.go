package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ItemInstanceService struct {
	Repo *repository.ItemInstanceRepository
}

func NewItemInstanceService(repo *repository.ItemInstanceRepository) *ItemInstanceService {
	return &ItemInstanceService{Repo: repo}
}

func (s *ItemInstanceService) GetByID(ctx context.Context, id string) (*model.ItemInstance, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ItemInstanceService) List(ctx context.Context, opts model.ListOptions) ([]model.ItemInstance, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ItemInstanceService) Create(ctx context.Context, loginID string, req *model.ItemInstanceRequest) (*model.ItemInstance, error, map[string]string) {
	m := &model.ItemInstance{}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	m.LocationId = req.LocationId
	if req.MaterialId != nil { m.MaterialId = *req.MaterialId }
	m.BatchNumber = req.BatchNumber
	m.SerialNumber = req.SerialNumber
	m.LotNumber = req.LotNumber
	m.ExpiryDate = req.ExpiryDate
	m.ManufacturedDate = req.ManufacturedDate
	m.Condition = req.Condition
	m.InspectionDueDate = req.InspectionDueDate
	if req.Quantity != nil { m.Quantity = *req.Quantity }
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ItemInstanceService) Update(ctx context.Context, loginID string, id string, req *model.ItemInstanceRequest) (*model.ItemInstance, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.LocationId != nil { m.LocationId = req.LocationId }
	if req.MaterialId != nil { m.MaterialId = *req.MaterialId }
	if req.BatchNumber != nil { m.BatchNumber = req.BatchNumber }
	if req.SerialNumber != nil { m.SerialNumber = req.SerialNumber }
	if req.LotNumber != nil { m.LotNumber = req.LotNumber }
	if req.ExpiryDate != nil { m.ExpiryDate = req.ExpiryDate }
	if req.ManufacturedDate != nil { m.ManufacturedDate = req.ManufacturedDate }
	if req.Condition != nil { m.Condition = req.Condition }
	if req.InspectionDueDate != nil { m.InspectionDueDate = req.InspectionDueDate }
	if req.Quantity != nil { m.Quantity = *req.Quantity }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ItemInstanceService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
