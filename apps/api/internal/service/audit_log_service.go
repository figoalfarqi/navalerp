package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type AuditLogService struct {
	Repo *repository.AuditLogRepository
}

func NewAuditLogService(repo *repository.AuditLogRepository) *AuditLogService {
	return &AuditLogService{Repo: repo}
}

func (s *AuditLogService) GetByID(ctx context.Context, id string) (*model.AuditLog, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *AuditLogService) List(ctx context.Context, opts model.ListOptions) ([]model.AuditLog, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *AuditLogService) Create(ctx context.Context, loginID string, req *model.AuditLogRequest) (*model.AuditLog, error, map[string]string) {
	m := &model.AuditLog{}
	m.UserId = req.UserId
	if req.Action != nil { m.Action = *req.Action }
	if req.EntityTable != nil { m.EntityTable = *req.EntityTable }
	m.EntityId = req.EntityId
	if req.OldValues != nil { m.OldValues = req.OldValues }
	if req.NewValues != nil { m.NewValues = req.NewValues }
	m.IpAddress = req.IpAddress
	m.UserAgent = req.UserAgent
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *AuditLogService) Update(ctx context.Context, loginID string, id string, req *model.AuditLogRequest) (*model.AuditLog, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.UserId != nil { m.UserId = req.UserId }
	if req.Action != nil { m.Action = *req.Action }
	if req.EntityTable != nil { m.EntityTable = *req.EntityTable }
	if req.EntityId != nil { m.EntityId = req.EntityId }
	if req.OldValues != nil { m.OldValues = req.OldValues }
	if req.NewValues != nil { m.NewValues = req.NewValues }
	if req.IpAddress != nil { m.IpAddress = req.IpAddress }
	if req.UserAgent != nil { m.UserAgent = req.UserAgent }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *AuditLogService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
