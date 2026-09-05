package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type RequisitionService struct {
	Repo *repository.RequisitionRepository
}

func NewRequisitionService(repo *repository.RequisitionRepository) *RequisitionService {
	return &RequisitionService{Repo: repo}
}

func (s *RequisitionService) GetByID(ctx context.Context, id string) (*model.Requisition, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *RequisitionService) List(ctx context.Context, opts model.ListOptions) ([]model.Requisition, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *RequisitionService) Create(ctx context.Context, loginID string, req *model.RequisitionRequest) (*model.Requisition, error, map[string]string) {
	m := &model.Requisition{}
	if req.RequisitionNumber != nil { m.RequisitionNumber = *req.RequisitionNumber }
	if req.OriginUnitId != nil { m.OriginUnitId = *req.OriginUnitId }
	m.WorkOrderId = req.WorkOrderId
	if req.Priority != nil { m.Priority = *req.Priority }
	if req.RequestedDate != nil { m.RequestedDate = *req.RequestedDate }
	m.RequiredByDate = req.RequiredByDate
	m.ApprovalStatus = req.ApprovalStatus
	m.ApprovedByUserId = req.ApprovedByUserId
	m.ApprovedAt = req.ApprovedAt
	m.TotalEstimatedCost = req.TotalEstimatedCost
	m.Justification = req.Justification
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *RequisitionService) Update(ctx context.Context, loginID string, id string, req *model.RequisitionRequest) (*model.Requisition, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.RequisitionNumber != nil { m.RequisitionNumber = *req.RequisitionNumber }
	if req.OriginUnitId != nil { m.OriginUnitId = *req.OriginUnitId }
	if req.WorkOrderId != nil { m.WorkOrderId = req.WorkOrderId }
	if req.Priority != nil { m.Priority = *req.Priority }
	if req.RequestedDate != nil { m.RequestedDate = *req.RequestedDate }
	if req.RequiredByDate != nil { m.RequiredByDate = req.RequiredByDate }
	if req.ApprovalStatus != nil { m.ApprovalStatus = req.ApprovalStatus }
	if req.ApprovedByUserId != nil { m.ApprovedByUserId = req.ApprovedByUserId }
	if req.ApprovedAt != nil { m.ApprovedAt = req.ApprovedAt }
	if req.TotalEstimatedCost != nil { m.TotalEstimatedCost = req.TotalEstimatedCost }
	if req.Justification != nil { m.Justification = req.Justification }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *RequisitionService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
