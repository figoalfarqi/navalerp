package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type StockAdjustmentService struct {
	Repo *repository.StockAdjustmentRepository
}

func NewStockAdjustmentService(repo *repository.StockAdjustmentRepository) *StockAdjustmentService {
	return &StockAdjustmentService{Repo: repo}
}

func (s *StockAdjustmentService) GetByID(ctx context.Context, id string) (*model.StockAdjustment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *StockAdjustmentService) List(ctx context.Context, opts model.ListOptions) ([]model.StockAdjustment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *StockAdjustmentService) Create(ctx context.Context, loginID string, req *model.StockAdjustmentRequest) (*model.StockAdjustment, error, map[string]string) {
	m := &model.StockAdjustment{}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.AdjustmentNumber != nil { m.AdjustmentNumber = *req.AdjustmentNumber }
	if req.AdjustmentDate != nil { m.AdjustmentDate = *req.AdjustmentDate }
	m.ConductedByUserId = req.ConductedByUserId
	if req.Reason != nil { m.Reason = *req.Reason }
	m.Status = req.Status
	m.Remarks = req.Remarks
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockAdjustmentService) Update(ctx context.Context, loginID string, id string, req *model.StockAdjustmentRequest) (*model.StockAdjustment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.AdjustmentNumber != nil { m.AdjustmentNumber = *req.AdjustmentNumber }
	if req.AdjustmentDate != nil { m.AdjustmentDate = *req.AdjustmentDate }
	if req.ConductedByUserId != nil { m.ConductedByUserId = req.ConductedByUserId }
	if req.Reason != nil { m.Reason = *req.Reason }
	if req.Status != nil { m.Status = req.Status }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockAdjustmentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
