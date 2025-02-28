package repository

import (
	"context"
	"testing"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	storage := config.NewStorage()
	repo := NewUserRepository(storage)

	user := newUser()

	err := repo.CreateUser(context.Background(), &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	dbUser := new(models.User)
	storage.Db.First(dbUser, user.ID)

	assert.Equal(t, dbUser.ID, user.ID)

	storage.Db.Where("1 = 1").Delete(&models.User{})
}

func TestGetUserByEmail(t *testing.T) {
	storage := config.NewStorage()
	repo := NewUserRepository(storage)

	user := newUser()
	storage.Db.Create(&user)

	dbUser, err := repo.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		t.Fatalf("Failed to get user by email: %v", err)
	}

	assert.Equal(t, dbUser.ID, user.ID)

	storage.Db.Where("1 = 1").Delete(&models.User{})
}
