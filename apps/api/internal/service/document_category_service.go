package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type DocumentCategoryService struct {
	Repo *repository.DocumentCategoryRepository
}

func NewDocumentCategoryService(repo *repository.DocumentCategoryRepository) *DocumentCategoryService {
	return &DocumentCategoryService{Repo: repo}
}

func (s *DocumentCategoryService) GetByID(ctx context.Context, id string) (*model.DocumentCategory, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *DocumentCategoryService) List(ctx context.Context, opts model.ListOptions) ([]model.DocumentCategory, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *DocumentCategoryService) Create(ctx context.Context, loginID string, req *model.DocumentCategoryRequest) (*model.DocumentCategory, error, map[string]string) {
	m := &model.DocumentCategory{}
	if req.CategoryCode != nil { m.CategoryCode = *req.CategoryCode }
	if req.CategoryName != nil { m.CategoryName = *req.CategoryName }
	m.RetentionYears = req.RetentionYears
	m.ConfidentialityLevel = req.ConfidentialityLevel
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DocumentCategoryService) Update(ctx context.Context, loginID string, id string, req *model.DocumentCategoryRequest) (*model.DocumentCategory, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.CategoryCode != nil { m.CategoryCode = *req.CategoryCode }
	if req.CategoryName != nil { m.CategoryName = *req.CategoryName }
	if req.RetentionYears != nil { m.RetentionYears = req.RetentionYears }
	if req.ConfidentialityLevel != nil { m.ConfidentialityLevel = req.ConfidentialityLevel }
	if req.Description != nil { m.Description = req.Description }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DocumentCategoryService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
