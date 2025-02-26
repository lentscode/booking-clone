package repository

import (
	"context"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/models"
)

type HostRepository struct {
	storage *config.Storage
}

func NewHostRepository(storage *config.Storage) *HostRepository {
	return &HostRepository{storage: storage}
}

func (r HostRepository) GetHost(ctx context.Context, id int64) (*models.Host, error) {
	host := new(models.Host)

	result := r.storage.Db.WithContext(ctx).First(host, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return host, nil
}

func (r HostRepository) GetHosts(ctx context.Context) ([]models.Host, error) {
	var hosts []models.Host

	result := r.storage.Db.WithContext(ctx).Find(&hosts)

	if result.Error != nil {
		return nil, result.Error
	}

	return hosts, nil
}
