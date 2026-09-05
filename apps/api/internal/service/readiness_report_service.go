package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ReadinessReportService struct {
	Repo *repository.ReadinessReportRepository
}

func NewReadinessReportService(repo *repository.ReadinessReportRepository) *ReadinessReportService {
	return &ReadinessReportService{Repo: repo}
}

func (s *ReadinessReportService) GetByID(ctx context.Context, id string) (*model.ReadinessReport, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ReadinessReportService) List(ctx context.Context, opts model.ListOptions) ([]model.ReadinessReport, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ReadinessReportService) Create(ctx context.Context, loginID string, req *model.ReadinessReportRequest) (*model.ReadinessReport, error, map[string]string) {
	m := &model.ReadinessReport{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.SnapshotTimestamp != nil { m.SnapshotTimestamp = *req.SnapshotTimestamp }
	if req.ReadinessCategory != nil { m.ReadinessCategory = *req.ReadinessCategory }
	if req.MroReadinessScore != nil { m.MroReadinessScore = *req.MroReadinessScore }
	if req.PersonnelManningScore != nil { m.PersonnelManningScore = *req.PersonnelManningScore }
	if req.LogisticsSupplyScore != nil { m.LogisticsSupplyScore = *req.LogisticsSupplyScore }
	m.Remarks = req.Remarks
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ReadinessReportService) Update(ctx context.Context, loginID string, id string, req *model.ReadinessReportRequest) (*model.ReadinessReport, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.SnapshotTimestamp != nil { m.SnapshotTimestamp = *req.SnapshotTimestamp }
	if req.ReadinessCategory != nil { m.ReadinessCategory = *req.ReadinessCategory }
	if req.MroReadinessScore != nil { m.MroReadinessScore = *req.MroReadinessScore }
	if req.PersonnelManningScore != nil { m.PersonnelManningScore = *req.PersonnelManningScore }
	if req.LogisticsSupplyScore != nil { m.LogisticsSupplyScore = *req.LogisticsSupplyScore }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ReadinessReportService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
