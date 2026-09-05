package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type FailureReportService struct {
	Repo *repository.FailureReportRepository
}

func NewFailureReportService(repo *repository.FailureReportRepository) *FailureReportService {
	return &FailureReportService{Repo: repo}
}

func (s *FailureReportService) GetByID(ctx context.Context, id string) (*model.FailureReport, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *FailureReportService) List(ctx context.Context, opts model.ListOptions) ([]model.FailureReport, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *FailureReportService) Create(ctx context.Context, loginID string, req *model.FailureReportRequest) (*model.FailureReport, error, map[string]string) {
	m := &model.FailureReport{}
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.ReportedByUserId != nil { m.ReportedByUserId = *req.ReportedByUserId }
	if req.ReportNumber != nil { m.ReportNumber = *req.ReportNumber }
	if req.IncidentDate != nil { m.IncidentDate = *req.IncidentDate }
	if req.Severity != nil { m.Severity = *req.Severity }
	m.FailureMode = req.FailureMode
	if req.Description != nil { m.Description = *req.Description }
	m.OperationalImpact = req.OperationalImpact
	m.ImmediateActionTaken = req.ImmediateActionTaken
	m.Status = req.Status
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *FailureReportService) Update(ctx context.Context, loginID string, id string, req *model.FailureReportRequest) (*model.FailureReport, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.ReportedByUserId != nil { m.ReportedByUserId = *req.ReportedByUserId }
	if req.ReportNumber != nil { m.ReportNumber = *req.ReportNumber }
	if req.IncidentDate != nil { m.IncidentDate = *req.IncidentDate }
	if req.Severity != nil { m.Severity = *req.Severity }
	if req.FailureMode != nil { m.FailureMode = req.FailureMode }
	if req.Description != nil { m.Description = *req.Description }
	if req.OperationalImpact != nil { m.OperationalImpact = req.OperationalImpact }
	if req.ImmediateActionTaken != nil { m.ImmediateActionTaken = req.ImmediateActionTaken }
	if req.Status != nil { m.Status = req.Status }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *FailureReportService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
