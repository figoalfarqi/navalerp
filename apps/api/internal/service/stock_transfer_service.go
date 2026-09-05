package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type StockTransferService struct {
	Repo *repository.StockTransferRepository
}

func NewStockTransferService(repo *repository.StockTransferRepository) *StockTransferService {
	return &StockTransferService{Repo: repo}
}

func (s *StockTransferService) GetByID(ctx context.Context, id string) (*model.StockTransfer, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *StockTransferService) List(ctx context.Context, opts model.ListOptions) ([]model.StockTransfer, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *StockTransferService) Create(ctx context.Context, loginID string, req *model.StockTransferRequest) (*model.StockTransfer, error, map[string]string) {
	m := &model.StockTransfer{}
	if req.TransferNumber != nil { m.TransferNumber = *req.TransferNumber }
	if req.FromWarehouseId != nil { m.FromWarehouseId = *req.FromWarehouseId }
	if req.ToWarehouseId != nil { m.ToWarehouseId = *req.ToWarehouseId }
	if req.MovementType != nil { m.MovementType = *req.MovementType }
	m.ScheduledDeparture = req.ScheduledDeparture
	m.ActualDeparture = req.ActualDeparture
	m.ScheduledArrival = req.ScheduledArrival
	m.ActualArrival = req.ActualArrival
	m.TransporterUnit = req.TransporterUnit
	m.Status = req.Status
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockTransferService) Update(ctx context.Context, loginID string, id string, req *model.StockTransferRequest) (*model.StockTransfer, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.TransferNumber != nil { m.TransferNumber = *req.TransferNumber }
	if req.FromWarehouseId != nil { m.FromWarehouseId = *req.FromWarehouseId }
	if req.ToWarehouseId != nil { m.ToWarehouseId = *req.ToWarehouseId }
	if req.MovementType != nil { m.MovementType = *req.MovementType }
	if req.ScheduledDeparture != nil { m.ScheduledDeparture = req.ScheduledDeparture }
	if req.ActualDeparture != nil { m.ActualDeparture = req.ActualDeparture }
	if req.ScheduledArrival != nil { m.ScheduledArrival = req.ScheduledArrival }
	if req.ActualArrival != nil { m.ActualArrival = req.ActualArrival }
	if req.TransporterUnit != nil { m.TransporterUnit = req.TransporterUnit }
	if req.Status != nil { m.Status = req.Status }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *StockTransferService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
