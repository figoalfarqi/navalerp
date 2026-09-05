package service

import (
	"context"
	"errors"
	"strings"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectTransportPhotoService struct {
	Repo        *repository.ProjectTransportPhotoRepository
	ProjectRepo *repository.ProjectRepository
	File        *FileUploadService
}

func NewProjectTransportPhotoService(
	repo *repository.ProjectTransportPhotoRepository,
	projectRepo *repository.ProjectRepository,
	file *FileUploadService,
) *ProjectTransportPhotoService {
	return &ProjectTransportPhotoService{Repo: repo, ProjectRepo: projectRepo, File: file}
}

func (s *ProjectTransportPhotoService) ensureAccess(ctx context.Context, checkerID *int, statusID int) error {
	if checkerID == nil {
		return nil
	}
	projectID, err := s.Repo.ResolveProjectIDByStatus(ctx, statusID)
	if err != nil {
		return err
	}
	allowed, err := s.ProjectRepo.CheckerCanAccess(ctx, projectID, *checkerID)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("checker is not assigned to this project")
	}
	return nil
}

func (s *ProjectTransportPhotoService) Create(ctx context.Context, userID int, checkerID *int, req *model.ProjectTransportPhotoRequest) (*model.ProjectTransportPhoto, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.ensureAccess(ctx, checkerID, req.ProjectTransportStatusID); err != nil {
		return nil, err, nil
	}
	if strings.Contains(req.PhotoURL, "/uploads/temp/") {
		if err := s.File.FinalizeManyFiles(map[string]*string{"project_transport_photo": &req.PhotoURL}); err != nil {
			return nil, err, nil
		}
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, checkerID)
	return item, err, nil
}
func (s *ProjectTransportPhotoService) Update(ctx context.Context, userID, id int, req *model.ProjectTransportPhotoRequest) (*model.ProjectTransportPhoto, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	old, err := s.Repo.GetByID(ctx, id, nil)
	if err != nil {
		return nil, err, nil
	}
	oldURL := old.PhotoURL
	newURL := &req.PhotoURL
	if oldURL != req.PhotoURL {
		if err := s.File.UpdateFile(&oldURL, &newURL); err != nil {
			return nil, err, nil
		}
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id, nil)
	return item, err, nil
}
func (s *ProjectTransportPhotoService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectTransportPhotoService) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransportPhoto, error) {
	return s.Repo.GetByID(ctx, id, checkerID)
}
func (s *ProjectTransportPhotoService) List(ctx context.Context, opts model.ListOptions, statusID *int) ([]model.ProjectTransportPhoto, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts), statusID)
}
