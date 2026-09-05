package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type GoodsReceiptService struct {
	Repo *repository.GoodsReceiptRepository
}

func NewGoodsReceiptService(repo *repository.GoodsReceiptRepository) *GoodsReceiptService {
	return &GoodsReceiptService{Repo: repo}
}

func (s *GoodsReceiptService) GetByID(ctx context.Context, id string) (*model.GoodsReceipt, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *GoodsReceiptService) List(ctx context.Context, opts model.ListOptions) ([]model.GoodsReceipt, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *GoodsReceiptService) Create(ctx context.Context, loginID string, req *model.GoodsReceiptRequest) (*model.GoodsReceipt, error, map[string]string) {
	m := &model.GoodsReceipt{}
	if req.ReceiptNumber != nil { m.ReceiptNumber = *req.ReceiptNumber }
	if req.PoId != nil { m.PoId = *req.PoId }
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.ReceivedDate != nil { m.ReceivedDate = *req.ReceivedDate }
	m.DeliveryOrderNumber = req.DeliveryOrderNumber
	if req.InspectedByUserId != nil { m.InspectedByUserId = *req.InspectedByUserId }
	m.InspectionPassed = req.InspectionPassed
	m.Remarks = req.Remarks
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *GoodsReceiptService) Update(ctx context.Context, loginID string, id string, req *model.GoodsReceiptRequest) (*model.GoodsReceipt, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ReceiptNumber != nil { m.ReceiptNumber = *req.ReceiptNumber }
	if req.PoId != nil { m.PoId = *req.PoId }
	if req.WarehouseId != nil { m.WarehouseId = *req.WarehouseId }
	if req.ReceivedDate != nil { m.ReceivedDate = *req.ReceivedDate }
	if req.DeliveryOrderNumber != nil { m.DeliveryOrderNumber = req.DeliveryOrderNumber }
	if req.InspectedByUserId != nil { m.InspectedByUserId = *req.InspectedByUserId }
	if req.InspectionPassed != nil { m.InspectionPassed = req.InspectionPassed }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *GoodsReceiptService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
