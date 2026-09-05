package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type JournalEntryService struct {
	Repo *repository.JournalEntryRepository
}

func NewJournalEntryService(repo *repository.JournalEntryRepository) *JournalEntryService {
	return &JournalEntryService{Repo: repo}
}

func (s *JournalEntryService) GetByID(ctx context.Context, id string) (*model.JournalEntry, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *JournalEntryService) List(ctx context.Context, opts model.ListOptions) ([]model.JournalEntry, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *JournalEntryService) Create(ctx context.Context, loginID string, req *model.JournalEntryRequest) (*model.JournalEntry, error, map[string]string) {
	m := &model.JournalEntry{}
	if req.EntryNumber != nil { m.EntryNumber = *req.EntryNumber }
	if req.EntryDate != nil { m.EntryDate = *req.EntryDate }
	if req.Description != nil { m.Description = *req.Description }
	if req.SourceModule != nil { m.SourceModule = *req.SourceModule }
	m.SourceReferenceId = req.SourceReferenceId
	m.IsPosted = req.IsPosted
	m.Lines = req.Lines
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *JournalEntryService) Update(ctx context.Context, loginID string, id string, req *model.JournalEntryRequest) (*model.JournalEntry, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.EntryNumber != nil { m.EntryNumber = *req.EntryNumber }
	if req.EntryDate != nil { m.EntryDate = *req.EntryDate }
	if req.Description != nil { m.Description = *req.Description }
	if req.SourceModule != nil { m.SourceModule = *req.SourceModule }
	if req.SourceReferenceId != nil { m.SourceReferenceId = req.SourceReferenceId }
	if req.IsPosted != nil { m.IsPosted = req.IsPosted }
	if req.Lines != nil { m.Lines = req.Lines }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *JournalEntryService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
