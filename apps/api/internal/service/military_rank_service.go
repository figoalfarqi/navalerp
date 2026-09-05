package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type MilitaryRankService struct {
	Repo *repository.MilitaryRankRepository
}

func NewMilitaryRankService(repo *repository.MilitaryRankRepository) *MilitaryRankService {
	return &MilitaryRankService{Repo: repo}
}

func (s *MilitaryRankService) GetByID(ctx context.Context, id string) (*model.MilitaryRank, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *MilitaryRankService) List(ctx context.Context, opts model.ListOptions) ([]model.MilitaryRank, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *MilitaryRankService) Create(ctx context.Context, loginID string, req *model.MilitaryRankRequest) (*model.MilitaryRank, error, map[string]string) {
	m := &model.MilitaryRank{}
	if req.RankCode != nil { m.RankCode = *req.RankCode }
	if req.RankName != nil { m.RankName = *req.RankName }
	if req.RankCategory != nil { m.RankCategory = *req.RankCategory }
	m.NatoRankCode = req.NatoRankCode
	if req.SeniorityOrder != nil { m.SeniorityOrder = *req.SeniorityOrder }
	m.IsActive = req.IsActive
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MilitaryRankService) Update(ctx context.Context, loginID string, id string, req *model.MilitaryRankRequest) (*model.MilitaryRank, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.RankCode != nil { m.RankCode = *req.RankCode }
	if req.RankName != nil { m.RankName = *req.RankName }
	if req.RankCategory != nil { m.RankCategory = *req.RankCategory }
	if req.NatoRankCode != nil { m.NatoRankCode = req.NatoRankCode }
	if req.SeniorityOrder != nil { m.SeniorityOrder = *req.SeniorityOrder }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *MilitaryRankService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
