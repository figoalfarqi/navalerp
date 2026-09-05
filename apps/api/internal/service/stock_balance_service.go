package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type StockBalanceService struct {
	Repo *repository.StockBalanceRepository
}

func NewStockBalanceService(repo *repository.StockBalanceRepository) *StockBalanceService {
	return &StockBalanceService{Repo: repo}
}

func (s *StockBalanceService) GetByID(ctx context.Context, id string) (*model.StockBalance, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *StockBalanceService) List(ctx context.Context, opts model.ListOptions) ([]model.StockBalance, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *StockBalanceService) Create(ctx context.Context, loginID string, req *model.StockBalanceRequest) (*model.StockBalance, error, map[string]string) {
	m := &model.StockBalance{}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	m.LocationId = req.LocationId
	if req.MaterialId != nil { m.MaterialId = *req.MaterialId }
	if req.QuantityOnHand != nil { m.QuantityOnHand = *req.QuantityOnHand }
	if req.QuantityReserved != nil { m.QuantityReserved = *req.QuantityReserved }
	if req.QuantityInTransit != nil { m.QuantityInTransit = *req.QuantityInTransit }
	m.LastCountDate = req.LastCountDate
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockBalanceService) Update(ctx context.Context, loginID string, id string, req *model.StockBalanceRequest) (*model.StockBalance, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.LocationId != nil { m.LocationId = req.LocationId }
	if req.MaterialId != nil { m.MaterialId = *req.MaterialId }
	if req.QuantityOnHand != nil { m.QuantityOnHand = *req.QuantityOnHand }
	if req.QuantityReserved != nil { m.QuantityReserved = *req.QuantityReserved }
	if req.QuantityInTransit != nil { m.QuantityInTransit = *req.QuantityInTransit }
	if req.LastCountDate != nil { m.LastCountDate = req.LastCountDate }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockBalanceService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
