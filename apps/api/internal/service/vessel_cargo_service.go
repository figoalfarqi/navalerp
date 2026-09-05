package service

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type VesselCargoService struct {
	Repo *repository.VesselCargoRepository
}

func NewVesselCargoService(repo *repository.VesselCargoRepository) *VesselCargoService {
	return &VesselCargoService{Repo: repo}
}
func (s *VesselCargoService) Create(ctx context.Context, userID int, req *model.VesselCargoRequest) (*model.VesselCargo, error, map[string]string) {
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
func (s *VesselCargoService) Update(ctx context.Context, userID, id int, req *model.VesselCargoRequest) (*model.VesselCargo, error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *VesselCargoService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *VesselCargoService) GetByID(ctx context.Context, id int) (*model.VesselCargo, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *VesselCargoService) List(ctx context.Context, opts model.ListOptions) ([]model.VesselCargo, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
