package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type MaterialService struct {
	Repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	return &MaterialService{Repo: repo}
}

func (s *MaterialService) GetByID(ctx context.Context, id string) (*model.Material, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *MaterialService) List(ctx context.Context, opts model.ListOptions) ([]model.Material, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *MaterialService) Create(ctx context.Context, loginID string, req *model.MaterialRequest) (*model.Material, error, map[string]string) {
	m := &model.Material{}
	if req.MaterialCode != nil { m.MaterialCode = *req.MaterialCode }
	m.Nsn = req.Nsn
	m.PartNumber = req.PartNumber
	m.OemName = req.OemName
	if req.MaterialName != nil { m.MaterialName = *req.MaterialName }
	if req.Category != nil { m.Category = *req.Category }
	if req.Uom != nil { m.Uom = *req.Uom }
	m.WeightKg = req.WeightKg
	m.MinStockLevel = req.MinStockLevel
	m.MaxStockLevel = req.MaxStockLevel
	m.ReorderPoint = req.ReorderPoint
	m.SafetyStock = req.SafetyStock
	m.ShelfLifeDays = req.ShelfLifeDays
	m.IsControlledItem = req.IsControlledItem
	m.UnitPriceIdr = req.UnitPriceIdr
	m.EquipmentLinks = req.EquipmentLinks
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MaterialService) Update(ctx context.Context, loginID string, id string, req *model.MaterialRequest) (*model.Material, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.MaterialCode != nil { m.MaterialCode = *req.MaterialCode }
	if req.Nsn != nil { m.Nsn = req.Nsn }
	if req.PartNumber != nil { m.PartNumber = req.PartNumber }
	if req.OemName != nil { m.OemName = req.OemName }
	if req.MaterialName != nil { m.MaterialName = *req.MaterialName }
	if req.Category != nil { m.Category = *req.Category }
	if req.Uom != nil { m.Uom = *req.Uom }
	if req.WeightKg != nil { m.WeightKg = req.WeightKg }
	if req.MinStockLevel != nil { m.MinStockLevel = req.MinStockLevel }
	if req.MaxStockLevel != nil { m.MaxStockLevel = req.MaxStockLevel }
	if req.ReorderPoint != nil { m.ReorderPoint = req.ReorderPoint }
	if req.SafetyStock != nil { m.SafetyStock = req.SafetyStock }
	if req.ShelfLifeDays != nil { m.ShelfLifeDays = req.ShelfLifeDays }
	if req.IsControlledItem != nil { m.IsControlledItem = req.IsControlledItem }
	if req.UnitPriceIdr != nil { m.UnitPriceIdr = req.UnitPriceIdr }
	if req.EquipmentLinks != nil { m.EquipmentLinks = req.EquipmentLinks }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MaterialService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
