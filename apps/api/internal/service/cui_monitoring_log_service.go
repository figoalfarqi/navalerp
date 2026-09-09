package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type CuiMonitoringLogService struct {
	Repo *repository.CuiMonitoringLogRepository
}

func NewCuiMonitoringLogService(repo *repository.CuiMonitoringLogRepository) *CuiMonitoringLogService {
	return &CuiMonitoringLogService{Repo: repo}
}

func (s *CuiMonitoringLogService) GetByID(ctx context.Context, id string) (*model.CuiMonitoringLog, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *CuiMonitoringLogService) List(ctx context.Context, opts model.ListOptions) ([]model.CuiMonitoringLog, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *CuiMonitoringLogService) Create(ctx context.Context, loginID string, req *model.CuiMonitoringLogRequest) (*model.CuiMonitoringLog, error, map[string]string) {
	m := &model.CuiMonitoringLog{}
	if req.CuiAssetId != nil { m.CuiAssetId = *req.CuiAssetId }
	if req.SensorCode != nil { m.SensorCode = *req.SensorCode }
	if req.SensorType != nil { m.SensorType = *req.SensorType }
	if req.LogTime != nil { m.LogTime = *req.LogTime }
	if req.MetricValue != nil { m.MetricValue = *req.MetricValue }
	if req.MetricUnit != nil { m.MetricUnit = *req.MetricUnit }
	if req.Status != nil { m.Status = *req.Status }
	m.VesselProximityMmsi = req.VesselProximityMmsi
	m.AnomalyScore = req.AnomalyScore
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiMonitoringLogService) Update(ctx context.Context, loginID string, id string, req *model.CuiMonitoringLogRequest) (*model.CuiMonitoringLog, error, map[string]string) {
	m := &model.CuiMonitoringLog{}
	if req.CuiAssetId != nil { m.CuiAssetId = *req.CuiAssetId }
	if req.SensorCode != nil { m.SensorCode = *req.SensorCode }
	if req.SensorType != nil { m.SensorType = *req.SensorType }
	if req.LogTime != nil { m.LogTime = *req.LogTime }
	if req.MetricValue != nil { m.MetricValue = *req.MetricValue }
	if req.MetricUnit != nil { m.MetricUnit = *req.MetricUnit }
	if req.Status != nil { m.Status = *req.Status }
	m.VesselProximityMmsi = req.VesselProximityMmsi
	m.AnomalyScore = req.AnomalyScore
	m.Description = req.Description
	err := s.Repo.Update(ctx, id, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiMonitoringLogService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
