package repository

import (
	"context"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/models"
)

type UserRepository struct {
	storage *config.Storage
}

func NewUserRepository(storage *config.Storage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := r.storage.Db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
