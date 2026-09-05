package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type WorkOrderService struct {
	Repo *repository.WorkOrderRepository
}

func NewWorkOrderService(repo *repository.WorkOrderRepository) *WorkOrderService {
	return &WorkOrderService{Repo: repo}
}

func (s *WorkOrderService) GetByID(ctx context.Context, id string) (*model.WorkOrder, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *WorkOrderService) List(ctx context.Context, opts model.ListOptions) ([]model.WorkOrder, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *WorkOrderService) Create(ctx context.Context, loginID string, req *model.WorkOrderRequest) (*model.WorkOrder, error, map[string]string) {
	m := &model.WorkOrder{}
	m.FailureReportId = req.FailureReportId
	m.PmScheduleId = req.PmScheduleId
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.WorkOrderNumber != nil { m.WorkOrderNumber = *req.WorkOrderNumber }
	if req.WorkOrderType != nil { m.WorkOrderType = *req.WorkOrderType }
	if req.Priority != nil { m.Priority = *req.Priority }
	m.ScheduledStartDate = req.ScheduledStartDate
	m.ScheduledEndDate = req.ScheduledEndDate
	m.ActualStartDate = req.ActualStartDate
	m.ActualEndDate = req.ActualEndDate
	m.LeadEngineerUserId = req.LeadEngineerUserId
	m.AssignedFacility = req.AssignedFacility
	m.Status = req.Status
	m.TotalLaborHours = req.TotalLaborHours
	m.EstimatedCost = req.EstimatedCost
	m.ActualCost = req.ActualCost
	m.CompletionNotes = req.CompletionNotes
	m.Tasks = req.Tasks
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *WorkOrderService) Update(ctx context.Context, loginID string, id string, req *model.WorkOrderRequest) (*model.WorkOrder, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.FailureReportId != nil { m.FailureReportId = req.FailureReportId }
	if req.PmScheduleId != nil { m.PmScheduleId = req.PmScheduleId }
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.WorkOrderNumber != nil { m.WorkOrderNumber = *req.WorkOrderNumber }
	if req.WorkOrderType != nil { m.WorkOrderType = *req.WorkOrderType }
	if req.Priority != nil { m.Priority = *req.Priority }
	if req.ScheduledStartDate != nil { m.ScheduledStartDate = req.ScheduledStartDate }
	if req.ScheduledEndDate != nil { m.ScheduledEndDate = req.ScheduledEndDate }
	if req.ActualStartDate != nil { m.ActualStartDate = req.ActualStartDate }
	if req.ActualEndDate != nil { m.ActualEndDate = req.ActualEndDate }
	if req.LeadEngineerUserId != nil { m.LeadEngineerUserId = req.LeadEngineerUserId }
	if req.AssignedFacility != nil { m.AssignedFacility = req.AssignedFacility }
	if req.Status != nil { m.Status = req.Status }
	if req.TotalLaborHours != nil { m.TotalLaborHours = req.TotalLaborHours }
	if req.EstimatedCost != nil { m.EstimatedCost = req.EstimatedCost }
	if req.ActualCost != nil { m.ActualCost = req.ActualCost }
	if req.CompletionNotes != nil { m.CompletionNotes = req.CompletionNotes }
	if req.Tasks != nil { m.Tasks = req.Tasks }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *WorkOrderService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
