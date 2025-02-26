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

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := new(models.User)

	result := r.storage.Db.WithContext(ctx).Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetUserBySessionId(ctx context.Context, sessionId string) (*models.User, error) {
	user := new(models.User)
	session := new(models.UserSession)

	result := r.storage.Db.WithContext(ctx).Where("session_id = ?", sessionId).First(session)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.storage.Db.WithContext(ctx).First(user, session.UserID)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) CreateSession(ctx context.Context, session *models.UserSession) error {
	result := r.storage.Db.WithContext(ctx).Create(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}