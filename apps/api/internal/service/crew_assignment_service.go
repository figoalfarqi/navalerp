package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type CrewAssignmentService struct {
	Repo *repository.CrewAssignmentRepository
}

func NewCrewAssignmentService(repo *repository.CrewAssignmentRepository) *CrewAssignmentService {
	return &CrewAssignmentService{Repo: repo}
}

func (s *CrewAssignmentService) GetByID(ctx context.Context, id string) (*model.CrewAssignment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *CrewAssignmentService) List(ctx context.Context, opts model.ListOptions) ([]model.CrewAssignment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *CrewAssignmentService) Create(ctx context.Context, loginID string, req *model.CrewAssignmentRequest) (*model.CrewAssignment, error, map[string]string) {
	m := &model.CrewAssignment{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.PersonnelId != nil { m.PersonnelId = *req.PersonnelId }
	if req.CrewRole != nil { m.CrewRole = *req.CrewRole }
	if req.Department != nil { m.Department = *req.Department }
	m.WatchBillDuty = req.WatchBillDuty
	if req.AssignedDate != nil { m.AssignedDate = *req.AssignedDate }
	m.RelievedDate = req.RelievedDate
	m.IsActive = req.IsActive
	m.Allowances = req.Allowances
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CrewAssignmentService) Update(ctx context.Context, loginID string, id string, req *model.CrewAssignmentRequest) (*model.CrewAssignment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.PersonnelId != nil { m.PersonnelId = *req.PersonnelId }
	if req.CrewRole != nil { m.CrewRole = *req.CrewRole }
	if req.Department != nil { m.Department = *req.Department }
	if req.WatchBillDuty != nil { m.WatchBillDuty = req.WatchBillDuty }
	if req.AssignedDate != nil { m.AssignedDate = *req.AssignedDate }
	if req.RelievedDate != nil { m.RelievedDate = req.RelievedDate }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if req.Allowances != nil { m.Allowances = req.Allowances }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *CrewAssignmentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
