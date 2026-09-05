package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type QualificationService struct {
	Repo *repository.QualificationRepository
}

func NewQualificationService(repo *repository.QualificationRepository) *QualificationService {
	return &QualificationService{Repo: repo}
}

func (s *QualificationService) GetByID(ctx context.Context, id string) (*model.Qualification, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *QualificationService) List(ctx context.Context, opts model.ListOptions) ([]model.Qualification, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *QualificationService) Create(ctx context.Context, loginID string, req *model.QualificationRequest) (*model.Qualification, error, map[string]string) {
	m := &model.Qualification{}
	if req.QualificationCode != nil { m.QualificationCode = *req.QualificationCode }
	if req.QualificationName != nil { m.QualificationName = *req.QualificationName }
	if req.QualificationCategory != nil { m.QualificationCategory = *req.QualificationCategory }
	if req.IssuingInstitution != nil { m.IssuingInstitution = *req.IssuingInstitution }
	m.ValidityYears = req.ValidityYears
	m.Description = req.Description
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *QualificationService) Update(ctx context.Context, loginID string, id string, req *model.QualificationRequest) (*model.Qualification, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.QualificationCode != nil { m.QualificationCode = *req.QualificationCode }
	if req.QualificationName != nil { m.QualificationName = *req.QualificationName }
	if req.QualificationCategory != nil { m.QualificationCategory = *req.QualificationCategory }
	if req.IssuingInstitution != nil { m.IssuingInstitution = *req.IssuingInstitution }
	if req.ValidityYears != nil { m.ValidityYears = req.ValidityYears }
	if req.Description != nil { m.Description = req.Description }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *QualificationService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
