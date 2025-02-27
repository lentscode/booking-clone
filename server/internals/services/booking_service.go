package services

import (
	"context"
	"errors"

	"github.com/lentscode/booking-server/internals/models"
	"github.com/lentscode/booking-server/internals/repository"
)

type BookingService struct {
	bookingRepo *repository.BookingRepository
}

func NewBookingService(bookingRepo *repository.BookingRepository) *BookingService {
	return &BookingService{bookingRepo: bookingRepo}
}

func (s *BookingService) CreateBooking(ctx context.Context, booking *models.Booking) error {
	alreadyExistingBookings, _ := s.bookingRepo.GetBookingsOfHostBetween(ctx, booking.HostID, booking.CheckInDate, booking.CheckOutDate)
	if len(alreadyExistingBookings) > 0 {
		return errors.New("host is not available for the given date range")
	}

	return s.bookingRepo.CreateBooking(ctx, booking)
}

func (s *BookingService) GetBookingsOfUser(ctx context.Context, userId uint) ([]models.Booking, error) {
	return s.bookingRepo.GetBookingsOfUser(ctx, userId)
}

func (s *BookingService) GetBooking(ctx context.Context, bookingId uint) (*models.Booking, error) {
	return s.bookingRepo.GetBooking(ctx, bookingId)
}
