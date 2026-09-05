package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type DockingRecordService struct {
	Repo *repository.DockingRecordRepository
}

func NewDockingRecordService(repo *repository.DockingRecordRepository) *DockingRecordService {
	return &DockingRecordService{Repo: repo}
}

func (s *DockingRecordService) GetByID(ctx context.Context, id string) (*model.DockingRecord, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *DockingRecordService) List(ctx context.Context, opts model.ListOptions) ([]model.DockingRecord, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *DockingRecordService) Create(ctx context.Context, loginID string, req *model.DockingRecordRequest) (*model.DockingRecord, error, map[string]string) {
	m := &model.DockingRecord{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.ShipyardName != nil { m.ShipyardName = *req.ShipyardName }
	if req.DockingType != nil { m.DockingType = *req.DockingType }
	if req.EntryDate != nil { m.EntryDate = *req.EntryDate }
	if req.ScheduledExitDate != nil { m.ScheduledExitDate = *req.ScheduledExitDate }
	m.ActualExitDate = req.ActualExitDate
	m.SeaTrialPassed = req.SeaTrialPassed
	m.ClassificationSurveyor = req.ClassificationSurveyor
	m.CertificateNumber = req.CertificateNumber
	m.TotalDockingCost = req.TotalDockingCost
	m.DockingSummary = req.DockingSummary
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DockingRecordService) Update(ctx context.Context, loginID string, id string, req *model.DockingRecordRequest) (*model.DockingRecord, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.ShipyardName != nil { m.ShipyardName = *req.ShipyardName }
	if req.DockingType != nil { m.DockingType = *req.DockingType }
	if req.EntryDate != nil { m.EntryDate = *req.EntryDate }
	if req.ScheduledExitDate != nil { m.ScheduledExitDate = *req.ScheduledExitDate }
	if req.ActualExitDate != nil { m.ActualExitDate = req.ActualExitDate }
	if req.SeaTrialPassed != nil { m.SeaTrialPassed = req.SeaTrialPassed }
	if req.ClassificationSurveyor != nil { m.ClassificationSurveyor = req.ClassificationSurveyor }
	if req.CertificateNumber != nil { m.CertificateNumber = req.CertificateNumber }
	if req.TotalDockingCost != nil { m.TotalDockingCost = req.TotalDockingCost }
	if req.DockingSummary != nil { m.DockingSummary = req.DockingSummary }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DockingRecordService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
