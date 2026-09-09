package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type CuiInspectionService struct {
	Repo *repository.CuiInspectionRepository
}

func NewCuiInspectionService(repo *repository.CuiInspectionRepository) *CuiInspectionService {
	return &CuiInspectionService{Repo: repo}
}

func (s *CuiInspectionService) GetByID(ctx context.Context, id string) (*model.CuiInspection, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *CuiInspectionService) List(ctx context.Context, opts model.ListOptions) ([]model.CuiInspection, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *CuiInspectionService) Create(ctx context.Context, loginID string, req *model.CuiInspectionRequest) (*model.CuiInspection, error, map[string]string) {
	m := &model.CuiInspection{}
	if req.InspectionNumber != nil {
		m.InspectionNumber = *req.InspectionNumber
	}
	if req.CuiAssetId != nil {
		m.CuiAssetId = *req.CuiAssetId
	}
	if req.ShipId != nil && *req.ShipId != "" {
		m.ShipId = req.ShipId
	} else {
		m.ShipId = nil
	}
	if req.InspectionDate != nil {
		m.InspectionDate = *req.InspectionDate
	}
	if req.InspectorOfficerId != nil && *req.InspectorOfficerId != "" {
		m.InspectorOfficerId = req.InspectorOfficerId
	} else {
		m.InspectorOfficerId = nil
	}
	if req.Method != nil {
		m.Method = *req.Method
	}
	if req.ConditionRating != nil {
		m.ConditionRating = *req.ConditionRating
	}
	m.Findings = req.Findings
	if req.RemedialActionRequired != nil {
		m.RemedialActionRequired = *req.RemedialActionRequired
	}
	m.NextInspectionDate = req.NextInspectionDate
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiInspectionService) Update(ctx context.Context, loginID string, id string, req *model.CuiInspectionRequest) (*model.CuiInspection, error, map[string]string) {
	m := &model.CuiInspection{}
	if req.InspectionNumber != nil {
		m.InspectionNumber = *req.InspectionNumber
	}
	if req.CuiAssetId != nil {
		m.CuiAssetId = *req.CuiAssetId
	}
	if req.ShipId != nil && *req.ShipId != "" {
		m.ShipId = req.ShipId
	} else {
		m.ShipId = nil
	}
	if req.InspectionDate != nil {
		m.InspectionDate = *req.InspectionDate
	}
	if req.InspectorOfficerId != nil && *req.InspectorOfficerId != "" {
		m.InspectorOfficerId = req.InspectorOfficerId
	} else {
		m.InspectorOfficerId = nil
	}
	if req.Method != nil {
		m.Method = *req.Method
	}
	if req.ConditionRating != nil {
		m.ConditionRating = *req.ConditionRating
	}
	m.Findings = req.Findings
	if req.RemedialActionRequired != nil {
		m.RemedialActionRequired = *req.RemedialActionRequired
	}
	m.NextInspectionDate = req.NextInspectionDate
	err := s.Repo.Update(ctx, id, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CuiInspectionService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
