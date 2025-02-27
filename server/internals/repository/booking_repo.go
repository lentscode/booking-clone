package repository

import (
	"context"
	"time"

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

func (r BookingRepository) GetBookingsOfUser(ctx context.Context, userId uint) ([]models.Booking, error) {
	var bookings []models.Booking

	result := r.storage.Db.WithContext(ctx).Where(&models.Booking{UserID: userId}, "user_id").Find(&bookings)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}

func (r BookingRepository) GetBooking(ctx context.Context, bookingId uint) (*models.Booking, error) {
	booking := new(models.Booking)

	if result := r.storage.Db.WithContext(ctx).First(booking, bookingId); result.Error != nil {
		return nil, result.Error
	}

	return booking, nil
}

func (r BookingRepository) GetBookingsOfHost(ctx context.Context, hostId uint) ([]models.Booking, error) {
	var bookings []models.Booking

	result := r.storage.Db.WithContext(ctx).Where(&models.Booking{HostID: hostId}, "host_id").Find(&bookings)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}

func (r BookingRepository) GetBookingsOfHostBetween(ctx context.Context, hostId uint, start time.Time, end time.Time) ([]models.Booking, error) {
	bookings := make([]models.Booking, 0)

	result := r.storage.Db.WithContext(ctx).Where("host_id = ? AND start_time BETWEEN ? AND ?", hostId, start, end).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}