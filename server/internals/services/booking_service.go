package services

import (
	"context"

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
	return s.bookingRepo.CreateBooking(ctx, booking)
}

func (s *BookingService) GetBookingsOfUser(ctx context.Context, userId uint) ([]models.Booking, error) {
	return s.bookingRepo.GetBookingsOfUser(ctx, userId)
}
