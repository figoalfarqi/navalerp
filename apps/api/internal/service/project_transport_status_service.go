package service

import (
	"context"
	"errors"
	"strings"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectTransportStatusService struct {
	Repo          *repository.ProjectTransportStatusRepository
	TransportRepo *repository.ProjectTransportRepository
	ProjectRepo   *repository.ProjectRepository
}

func NewProjectTransportStatusService(
	repo *repository.ProjectTransportStatusRepository,
	transportRepo *repository.ProjectTransportRepository,
	projectRepo *repository.ProjectRepository,
) *ProjectTransportStatusService {
	return &ProjectTransportStatusService{Repo: repo, TransportRepo: transportRepo, ProjectRepo: projectRepo}
}

func (s *ProjectTransportStatusService) prepare(
	ctx context.Context,
	checkerID *int,
	excludedStatusID *int,
	req *model.ProjectTransportStatusRequest,
) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	if req.ProjectTransportStatusTypeID < 1 || req.ProjectTransportStatusTypeID > 8 {
		return errors.New("validation error"), map[string]string{
			"project_transport_status_type_id": "must be between 1 and 8",
		}
	}
	projectID, err := s.TransportRepo.ResolveProjectID(ctx, req.ProjectTransportID)
	if err != nil {
		return err, nil
	}
	if checkerID != nil {
		allowed, err := s.ProjectRepo.CheckerCanAccess(ctx, projectID, *checkerID)
		if err != nil {
			return err, nil
		}
		if !allowed {
			return errors.New("checker is not assigned to this project"), nil
		}
	}
	if req.ProjectTransportStatusTypeID == 5 {
		hasOrigin, err := s.Repo.HasStatusExcept(
			ctx,
			req.ProjectTransportID,
			1,
			excludedStatusID,
		)
		if err != nil {
			return err, nil
		}
		if !hasOrigin {
			fraud := 1
			req.IsFraud = &fraud
			note := ""
			if req.ProjectTransportStatusNote != nil {
				note = strings.TrimSpace(*req.ProjectTransportStatusNote)
			}
			if note == "" {
				note = "Arrived at destination without an origin arrival status"
			}
			req.ProjectTransportStatusNote = &note
		}
	}
	if req.IsFraud != nil && *req.IsFraud == 1 {
		note := ""
		if req.ProjectTransportStatusNote != nil {
			note = strings.TrimSpace(*req.ProjectTransportStatusNote)
		}
		if note == "" {
			note = "Status marked as fraud by checker"
			req.ProjectTransportStatusNote = &note
		}
	}
	return nil, nil
}

func (s *ProjectTransportStatusService) Create(ctx context.Context, userID int, checkerID *int, req *model.ProjectTransportStatusRequest) (*model.ProjectTransportStatus, error, map[string]string) {
	if err, fields := s.prepare(ctx, checkerID, nil, req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, checkerID)
	return item, err, nil
}
func (s *ProjectTransportStatusService) Update(ctx context.Context, userID, id int, req *model.ProjectTransportStatusRequest) (*model.ProjectTransportStatus, error, map[string]string) {
	if err, fields := s.prepare(ctx, nil, &id, req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectTransportStatusService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectTransportStatusService) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransportStatus, error) {
	return s.Repo.GetByID(ctx, id, checkerID)
}
func (s *ProjectTransportStatusService) List(ctx context.Context, opts model.ListOptions, transportID *int) ([]model.ProjectTransportStatus, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts), transportID)
}
