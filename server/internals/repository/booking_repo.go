package repository

import (
	"context"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/models"
)

type BookingRepository struct {
	storage *config.Storage
}

func NewBookingRepository(storage *config.Storage) *BookingRepository {
	return &BookingRepository{storage: storage}
}

func (r BookingRepository) CreateBooking(ctx context.Context, booking *models.Booking) error {
	result := r.storage.Db.WithContext(ctx).Create(booking)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r BookingRepository) GetBookingsOfUser(ctx context.Context, userId int64) ([]models.Booking, error) {
	var bookings []models.Booking

	result := r.storage.Db.WithContext(ctx).Where(&models.Booking{UserID: userId}, "user_id").Find(&bookings)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}
