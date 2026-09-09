package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type CuiAssetService struct {
	Repo *repository.CuiAssetRepository
}

func NewCuiAssetService(repo *repository.CuiAssetRepository) *CuiAssetService {
	return &CuiAssetService{Repo: repo}
}

func (s *CuiAssetService) GetByID(ctx context.Context, id string) (*model.CuiAsset, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *CuiAssetService) List(ctx context.Context, opts model.ListOptions) ([]model.CuiAsset, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *CuiAssetService) Create(ctx context.Context, loginID string, req *model.CuiAssetRequest) (*model.CuiAsset, error, map[string]string) {
	m := &model.CuiAsset{}
	if req.AssetCode != nil { m.AssetCode = *req.AssetCode }
	if req.AssetName != nil { m.AssetName = *req.AssetName }
	if req.AssetType != nil { m.AssetType = *req.AssetType }
	if req.OperatorName != nil { m.OperatorName = *req.OperatorName }
	m.TheaterId = req.TheaterId
	m.DepthMeters = req.DepthMeters
	m.LengthKm = req.LengthKm
	if req.Latitude != nil { m.Latitude = *req.Latitude }
	if req.Longitude != nil { m.Longitude = *req.Longitude }
	m.StartCoordinates = req.StartCoordinates
	m.EndCoordinates = req.EndCoordinates
	if req.Status != nil { m.Status = *req.Status }
	if req.HealthScore != nil { m.HealthScore = *req.HealthScore }
	if req.ProtectionPriority != nil { m.ProtectionPriority = *req.ProtectionPriority }
	m.LastInspectedAt = req.LastInspectedAt
	m.NextInspectionDue = req.NextInspectionDue
	m.Notes = req.Notes
	m.MonitoringLogs = req.MonitoringLogs
	m.Alerts = req.Alerts
	m.Inspections = req.Inspections
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiAssetService) Update(ctx context.Context, loginID string, id string, req *model.CuiAssetRequest) (*model.CuiAsset, error, map[string]string) {
	m := &model.CuiAsset{}
	if req.AssetCode != nil { m.AssetCode = *req.AssetCode }
	if req.AssetName != nil { m.AssetName = *req.AssetName }
	if req.AssetType != nil { m.AssetType = *req.AssetType }
	if req.OperatorName != nil { m.OperatorName = *req.OperatorName }
	m.TheaterId = req.TheaterId
	m.DepthMeters = req.DepthMeters
	m.LengthKm = req.LengthKm
	if req.Latitude != nil { m.Latitude = *req.Latitude }
	if req.Longitude != nil { m.Longitude = *req.Longitude }
	m.StartCoordinates = req.StartCoordinates
	m.EndCoordinates = req.EndCoordinates
	if req.Status != nil { m.Status = *req.Status }
	if req.HealthScore != nil { m.HealthScore = *req.HealthScore }
	if req.ProtectionPriority != nil { m.ProtectionPriority = *req.ProtectionPriority }
	m.LastInspectedAt = req.LastInspectedAt
	m.NextInspectionDue = req.NextInspectionDue
	m.Notes = req.Notes
	m.MonitoringLogs = req.MonitoringLogs
	m.Alerts = req.Alerts
	m.Inspections = req.Inspections
	err := s.Repo.Update(ctx, id, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiAssetService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
