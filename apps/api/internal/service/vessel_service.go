package service

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type VesselService struct{ Repo *repository.VesselRepository }

func NewVesselService(repo *repository.VesselRepository) *VesselService {
	return &VesselService{Repo: repo}
}
func (s *VesselService) Create(ctx context.Context, userID int, req *model.VesselRequest) (*model.Vessel, error, map[string]string) {
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
func (s *VesselService) Update(ctx context.Context, userID, id int, req *model.VesselRequest) (*model.Vessel, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *VesselService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *VesselService) GetByID(ctx context.Context, id int) (*model.Vessel, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *VesselService) List(ctx context.Context, opts model.ListOptions) ([]model.Vessel, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
