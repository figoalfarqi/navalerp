package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PmScheduleService struct {
	Repo *repository.PmScheduleRepository
}

func NewPmScheduleService(repo *repository.PmScheduleRepository) *PmScheduleService {
	return &PmScheduleService{Repo: repo}
}

func (s *PmScheduleService) GetByID(ctx context.Context, id string) (*model.PmSchedule, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *PmScheduleService) List(ctx context.Context, opts model.ListOptions) ([]model.PmSchedule, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *PmScheduleService) Create(ctx context.Context, loginID string, req *model.PmScheduleRequest) (*model.PmSchedule, error, map[string]string) {
	m := &model.PmSchedule{}
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.PmCode != nil { m.PmCode = *req.PmCode }
	if req.PmTitle != nil { m.PmTitle = *req.PmTitle }
	m.IntervalHours = req.IntervalHours
	m.IntervalDays = req.IntervalDays
	m.LastPerformedAt = req.LastPerformedAt
	m.NextDueAt = req.NextDueAt
	if req.TaskInstructions != nil { m.TaskInstructions = *req.TaskInstructions }
	m.EstimatedDurationHours = req.EstimatedDurationHours
	m.IsActive = req.IsActive
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PmScheduleService) Update(ctx context.Context, loginID string, id string, req *model.PmScheduleRequest) (*model.PmSchedule, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.EquipmentId != nil { m.EquipmentId = *req.EquipmentId }
	if req.PmCode != nil { m.PmCode = *req.PmCode }
	if req.PmTitle != nil { m.PmTitle = *req.PmTitle }
	if req.IntervalHours != nil { m.IntervalHours = req.IntervalHours }
	if req.IntervalDays != nil { m.IntervalDays = req.IntervalDays }
	if req.LastPerformedAt != nil { m.LastPerformedAt = req.LastPerformedAt }
	if req.NextDueAt != nil { m.NextDueAt = req.NextDueAt }
	if req.TaskInstructions != nil { m.TaskInstructions = *req.TaskInstructions }
	if req.EstimatedDurationHours != nil { m.EstimatedDurationHours = req.EstimatedDurationHours }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PmScheduleService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
