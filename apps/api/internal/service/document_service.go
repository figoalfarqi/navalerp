package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type DocumentService struct {
	Repo *repository.DocumentRepository
}

func NewDocumentService(repo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{Repo: repo}
}

func (s *DocumentService) GetByID(ctx context.Context, id string) (*model.Document, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *DocumentService) List(ctx context.Context, opts model.ListOptions) ([]model.Document, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *DocumentService) Create(ctx context.Context, loginID string, req *model.DocumentRequest) (*model.Document, error, map[string]string) {
	m := &model.Document{}
	if req.DocumentNumber != nil { m.DocumentNumber = *req.DocumentNumber }
	if req.Title != nil { m.Title = *req.Title }
	if req.CategoryId != nil { m.CategoryId = *req.CategoryId }
	m.OriginatingUnitId = req.OriginatingUnitId
	m.ClassificationLevel = req.ClassificationLevel
	if req.EffectiveDate != nil { m.EffectiveDate = *req.EffectiveDate }
	m.ExpiryDate = req.ExpiryDate
	m.Status = req.Status
	m.ApprovedByUserId = req.ApprovedByUserId
	m.Versions = req.Versions
	m.Links = req.Links
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DocumentService) Update(ctx context.Context, loginID string, id string, req *model.DocumentRequest) (*model.Document, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.DocumentNumber != nil { m.DocumentNumber = *req.DocumentNumber }
	if req.Title != nil { m.Title = *req.Title }
	if req.CategoryId != nil { m.CategoryId = *req.CategoryId }
	if req.OriginatingUnitId != nil { m.OriginatingUnitId = req.OriginatingUnitId }
	if req.ClassificationLevel != nil { m.ClassificationLevel = req.ClassificationLevel }
	if req.EffectiveDate != nil { m.EffectiveDate = *req.EffectiveDate }
	if req.ExpiryDate != nil { m.ExpiryDate = req.ExpiryDate }
	if req.Status != nil { m.Status = req.Status }
	if req.ApprovedByUserId != nil { m.ApprovedByUserId = req.ApprovedByUserId }
	if req.Versions != nil { m.Versions = req.Versions }
	if req.Links != nil { m.Links = req.Links }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *DocumentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
