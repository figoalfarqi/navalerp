package service

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectTruckAssignmentService struct {
	Repo *repository.ProjectTruckAssignmentRepository
}

func NewProjectTruckAssignmentService(repo *repository.ProjectTruckAssignmentRepository) *ProjectTruckAssignmentService {
	return &ProjectTruckAssignmentService{Repo: repo}
}
func (s *ProjectTruckAssignmentService) Create(ctx context.Context, userID int, req *model.ProjectTruckAssignmentRequest) (*model.ProjectTruckAssignment, error, map[string]string) {
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
func (s *ProjectTruckAssignmentService) Update(ctx context.Context, userID, id int, req *model.ProjectTruckAssignmentRequest) (*model.ProjectTruckAssignment, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *ProjectTruckAssignmentService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectTruckAssignmentService) GetByID(ctx context.Context, id int) (*model.ProjectTruckAssignment, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *ProjectTruckAssignmentService) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectTruckAssignment, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
