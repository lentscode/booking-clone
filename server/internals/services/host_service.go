package services

import (
	"context"

	"github.com/lentscode/booking-server/internals/models"
	"github.com/lentscode/booking-server/internals/repository"
)

type HostService struct {
	hostRepo *repository.HostRepository
}

func NewHostService(hostRepo *repository.HostRepository) *HostService {
	return &HostService{hostRepo: hostRepo}
}

func (s *HostService) GetHosts(ctx context.Context) ([]models.Host, error) {
	return s.hostRepo.GetHosts(ctx)
}

func (s *HostService) GetHost(ctx context.Context, id int64) (*models.Host, error) {
	return s.hostRepo.GetHost(ctx, id)
}
