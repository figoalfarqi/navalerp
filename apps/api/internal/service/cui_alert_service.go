package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type CuiAlertService struct {
	Repo *repository.CuiAlertRepository
}

func NewCuiAlertService(repo *repository.CuiAlertRepository) *CuiAlertService {
	return &CuiAlertService{Repo: repo}
}

func (s *CuiAlertService) GetByID(ctx context.Context, id string) (*model.CuiAlert, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *CuiAlertService) List(ctx context.Context, opts model.ListOptions) ([]model.CuiAlert, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *CuiAlertService) Create(ctx context.Context, loginID string, req *model.CuiAlertRequest) (*model.CuiAlert, error, map[string]string) {
	m := &model.CuiAlert{}
	if req.AlertCode != nil { m.AlertCode = *req.AlertCode }
	if req.CuiAssetId != nil { m.CuiAssetId = *req.CuiAssetId }
	if req.AlertType != nil { m.AlertType = *req.AlertType }
	if req.Severity != nil { m.Severity = *req.Severity }
	if req.DetectedAt != nil { m.DetectedAt = *req.DetectedAt }
	m.AssignedShipId = req.AssignedShipId
	if req.Status != nil { m.Status = *req.Status }
	if req.AiConfidence != nil { m.AiConfidence = *req.AiConfidence }
	m.RecommendedAction = req.RecommendedAction
	m.ResolutionNotes = req.ResolutionNotes
	m.ResolvedAt = req.ResolvedAt
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiAlertService) Update(ctx context.Context, loginID string, id string, req *model.CuiAlertRequest) (*model.CuiAlert, error, map[string]string) {
	m := &model.CuiAlert{}
	if req.AlertCode != nil { m.AlertCode = *req.AlertCode }
	if req.CuiAssetId != nil { m.CuiAssetId = *req.CuiAssetId }
	if req.AlertType != nil { m.AlertType = *req.AlertType }
	if req.Severity != nil { m.Severity = *req.Severity }
	if req.DetectedAt != nil { m.DetectedAt = *req.DetectedAt }
	m.AssignedShipId = req.AssignedShipId
	if req.Status != nil { m.Status = *req.Status }
	if req.AiConfidence != nil { m.AiConfidence = *req.AiConfidence }
	m.RecommendedAction = req.RecommendedAction
	m.ResolutionNotes = req.ResolutionNotes
	m.ResolvedAt = req.ResolvedAt
	err := s.Repo.Update(ctx, id, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiAlertService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
