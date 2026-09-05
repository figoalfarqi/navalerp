package service

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ProjectCheckerAssignmentService struct {
	Repo *repository.ProjectCheckerAssignmentRepository
}

func NewProjectCheckerAssignmentService(repo *repository.ProjectCheckerAssignmentRepository) *ProjectCheckerAssignmentService {
	return &ProjectCheckerAssignmentService{Repo: repo}
}
func (s *ProjectCheckerAssignmentService) Create(ctx context.Context, userID int, req *model.ProjectCheckerAssignmentRequest) (*model.ProjectCheckerAssignment, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *ProjectCheckerAssignmentService) Update(ctx context.Context, userID, id int, req *model.ProjectCheckerAssignmentRequest) (*model.ProjectCheckerAssignment, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *ProjectCheckerAssignmentService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectCheckerAssignmentService) GetByID(ctx context.Context, id int) (*model.ProjectCheckerAssignment, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *ProjectCheckerAssignmentService) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectCheckerAssignment, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
func (s *ProjectCheckerAssignmentService) Default(ctx context.Context, checkerID int) (*model.ProjectCheckerAssignment, error) {
	return s.Repo.GetDefaultForChecker(ctx, checkerID)
}
