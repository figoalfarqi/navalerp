package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PurchaseOrderService struct {
	Repo *repository.PurchaseOrderRepository
}

func NewPurchaseOrderService(repo *repository.PurchaseOrderRepository) *PurchaseOrderService {
	return &PurchaseOrderService{Repo: repo}
}

func (s *PurchaseOrderService) GetByID(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *PurchaseOrderService) List(ctx context.Context, opts model.ListOptions) ([]model.PurchaseOrder, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *PurchaseOrderService) Create(ctx context.Context, loginID string, req *model.PurchaseOrderRequest) (*model.PurchaseOrder, error, map[string]string) {
	m := &model.PurchaseOrder{}
	if req.PoNumber != nil { m.PoNumber = *req.PoNumber }
	m.ContractId = req.ContractId
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	if req.IssuingUnitId != nil { m.IssuingUnitId = *req.IssuingUnitId }
	if req.OrderDate != nil { m.OrderDate = *req.OrderDate }
	m.DeliveryDeadline = req.DeliveryDeadline
	m.DestinationWarehouseId = req.DestinationWarehouseId
	if req.TotalAmount != nil { m.TotalAmount = *req.TotalAmount }
	m.TaxAmount = req.TaxAmount
	m.Status = req.Status
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PurchaseOrderService) Update(ctx context.Context, loginID string, id string, req *model.PurchaseOrderRequest) (*model.PurchaseOrder, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.PoNumber != nil { m.PoNumber = *req.PoNumber }
	if req.ContractId != nil { m.ContractId = req.ContractId }
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	if req.IssuingUnitId != nil { m.IssuingUnitId = *req.IssuingUnitId }
	if req.OrderDate != nil { m.OrderDate = *req.OrderDate }
	if req.DeliveryDeadline != nil { m.DeliveryDeadline = req.DeliveryDeadline }
	if req.DestinationWarehouseId != nil { m.DestinationWarehouseId = req.DestinationWarehouseId }
	if req.TotalAmount != nil { m.TotalAmount = *req.TotalAmount }
	if req.TaxAmount != nil { m.TaxAmount = req.TaxAmount }
	if req.Status != nil { m.Status = req.Status }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PurchaseOrderService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
