package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ReadinessAlertService struct {
	Repo *repository.ReadinessAlertRepository
}

func NewReadinessAlertService(repo *repository.ReadinessAlertRepository) *ReadinessAlertService {
	return &ReadinessAlertService{Repo: repo}
}

func (s *ReadinessAlertService) GetByID(ctx context.Context, id string) (*model.ReadinessAlert, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ReadinessAlertService) List(ctx context.Context, opts model.ListOptions) ([]model.ReadinessAlert, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ReadinessAlertService) Create(ctx context.Context, loginID string, req *model.ReadinessAlertRequest) (*model.ReadinessAlert, error, map[string]string) {
	m := &model.ReadinessAlert{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	m.EquipmentId = req.EquipmentId
	if req.Severity != nil { m.Severity = *req.Severity }
	if req.AlertType != nil { m.AlertType = *req.AlertType }
	if req.AlertMessage != nil { m.AlertMessage = *req.AlertMessage }
	m.IsAcknowledged = req.IsAcknowledged
	m.AcknowledgedByUserId = req.AcknowledgedByUserId
	m.AcknowledgedAt = req.AcknowledgedAt
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ReadinessAlertService) Update(ctx context.Context, loginID string, id string, req *model.ReadinessAlertRequest) (*model.ReadinessAlert, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.EquipmentId != nil { m.EquipmentId = req.EquipmentId }
	if req.Severity != nil { m.Severity = *req.Severity }
	if req.AlertType != nil { m.AlertType = *req.AlertType }
	if req.AlertMessage != nil { m.AlertMessage = *req.AlertMessage }
	if req.IsAcknowledged != nil { m.IsAcknowledged = req.IsAcknowledged }
	if req.AcknowledgedByUserId != nil { m.AcknowledgedByUserId = req.AcknowledgedByUserId }
	if req.AcknowledgedAt != nil { m.AcknowledgedAt = req.AcknowledgedAt }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ReadinessAlertService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
