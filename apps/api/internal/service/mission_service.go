package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type MissionService struct {
	Repo *repository.MissionRepository
}

func NewMissionService(repo *repository.MissionRepository) *MissionService {
	return &MissionService{Repo: repo}
}

func (s *MissionService) GetByID(ctx context.Context, id string) (*model.Mission, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *MissionService) List(ctx context.Context, opts model.ListOptions) ([]model.Mission, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *MissionService) Create(ctx context.Context, loginID string, req *model.MissionRequest) (*model.Mission, error, map[string]string) {
	m := &model.Mission{}
	if req.TheaterId != nil { m.TheaterId = *req.TheaterId }
	if req.MissionCode != nil { m.MissionCode = *req.MissionCode }
	if req.MissionName != nil { m.MissionName = *req.MissionName }
	if req.MissionType != nil { m.MissionType = *req.MissionType }
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	m.EndDate = req.EndDate
	m.CommandingOfficerUserId = req.CommandingOfficerUserId
	m.MissionStatus = req.MissionStatus
	m.AssignedShips = req.AssignedShips
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MissionService) Update(ctx context.Context, loginID string, id string, req *model.MissionRequest) (*model.Mission, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.TheaterId != nil { m.TheaterId = *req.TheaterId }
	if req.MissionCode != nil { m.MissionCode = *req.MissionCode }
	if req.MissionName != nil { m.MissionName = *req.MissionName }
	if req.MissionType != nil { m.MissionType = *req.MissionType }
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	if req.EndDate != nil { m.EndDate = req.EndDate }
	if req.CommandingOfficerUserId != nil { m.CommandingOfficerUserId = req.CommandingOfficerUserId }
	if req.MissionStatus != nil { m.MissionStatus = req.MissionStatus }
	if req.AssignedShips != nil { m.AssignedShips = req.AssignedShips }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MissionService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
