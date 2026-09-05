package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectTransportService struct {
	Repo           *repository.ProjectTransportRepository
	ProjectRepo    *repository.ProjectRepository
	AssignmentRepo *repository.ProjectTruckAssignmentRepository
}

func NewProjectTransportService(
	repo *repository.ProjectTransportRepository,
	projectRepo *repository.ProjectRepository,
	assignmentRepo *repository.ProjectTruckAssignmentRepository,
) *ProjectTransportService {
	return &ProjectTransportService{Repo: repo, ProjectRepo: projectRepo, AssignmentRepo: assignmentRepo}
}

func (s *ProjectTransportService) validateAccessAndTruck(
	ctx context.Context,
	checkerID *int,
	req *model.ProjectTransportRequest,
) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	if checkerID != nil {
		allowed, err := s.ProjectRepo.CheckerCanAccess(ctx, req.ProjectID, *checkerID)
		if err != nil {
			return err, nil
		}
		if !allowed {
			return errors.New("checker is not assigned to this project"), nil
		}
	}
	if req.TruckID == nil || req.ProjectTruckAssignmentID == nil {
		return errors.New("truck_id and project_truck_assignment_id are required together"), map[string]string{
			"truck_id":                    "active project truck assignment is required",
			"project_truck_assignment_id": "active project truck assignment is required",
		}
	}
	valid, err := s.AssignmentRepo.IsValid(ctx, *req.ProjectTruckAssignmentID, req.ProjectID, *req.TruckID)
	if err != nil {
		return err, nil
	}
	if !valid {
		return errors.New("project truck assignment is inactive or outside its access period"), nil
	}
	return nil, nil
}

func (s *ProjectTransportService) Create(
	ctx context.Context,
	userID int,
	checkerID *int,
	req *model.ProjectTransportRequest,
) (*model.ProjectTransport, error, map[string]string) {
	if err, fields := s.validateAccessAndTruck(ctx, checkerID, req); err != nil {
		return nil, err, fields
	}
	// A checker selects an assigned truck; its driver/vendor snapshot must come
	// from that truck and cannot be overridden by a client-supplied payload.
	if checkerID != nil {
		req.TransportNumber = nil
		req.DriverID = nil
		req.TransportVendorID = nil
		req.IsCompleted = nil
	}
	number := ""
	if req.TransportNumber != nil {
		number = strings.TrimSpace(*req.TransportNumber)
	}
	if number == "" {
		number = fmt.Sprintf("TR-%d-%d", req.ProjectID, time.Now().UnixNano())
	}
	id, err := s.Repo.Create(ctx, req, number, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, checkerID)
	return item, err, nil
}

func (s *ProjectTransportService) Update(ctx context.Context, userID, id int, req *model.ProjectTransportRequest) (*model.ProjectTransport, error, map[string]string) {
	if err, fields := s.validateAccessAndTruck(ctx, nil, req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectTransportService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectTransportService) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransport, error) {
	return s.Repo.GetByID(ctx, id, checkerID)
}
func (s *ProjectTransportService) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectTransport, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
func (s *ProjectTransportService) ListOperational(ctx context.Context, opts model.ListOptions) ([]model.DriverProjectTransport, error) {
	return s.Repo.ListOperational(ctx, normalizeListOptions(opts))
}
