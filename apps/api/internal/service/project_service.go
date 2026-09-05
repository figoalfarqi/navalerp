package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectService struct{ Repo *repository.ProjectRepository }

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) validate(req *model.ProjectRequest) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	fields := map[string]string{}
	if err := validateDate(req.StartDate); err != nil {
		fields["start_date"] = "must use YYYY-MM-DD"
	}
	if err := validateDate(req.EndDate); err != nil {
		fields["end_date"] = "must use YYYY-MM-DD"
	}
	if len(fields) > 0 {
		return errors.New("validation error"), fields
	}
	return nil, nil
}

func (s *ProjectService) Create(ctx context.Context, userID int, req *model.ProjectRequest) (*model.Project, error, map[string]string) {
	if err, fields := s.validate(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}

func (s *ProjectService) Update(ctx context.Context, userID, id int, req *model.ProjectRequest) (*model.Project, error, map[string]string) {
	if err, fields := s.validate(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectService) GetByID(ctx context.Context, id int, checkerID *int) (*model.Project, error) {
	return s.Repo.GetByID(ctx, id, checkerID)
}
func (s *ProjectService) List(ctx context.Context, opts model.ListOptions) ([]model.Project, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
