package services

import (
	"context"
	"sort"
	"time"

	"github.com/lentscode/booking-server/internals/models"
	"github.com/lentscode/booking-server/internals/repository"
)

type HostService struct {
	hostRepo    *repository.HostRepository
	bookingRepo *repository.BookingRepository
}

func NewHostService(hostRepo *repository.HostRepository, bookingRepo *repository.BookingRepository) *HostService {
	return &HostService{hostRepo: hostRepo, bookingRepo: bookingRepo}
}

func (s *HostService) GetHosts(ctx context.Context) ([]models.Host, error) {
	return s.hostRepo.GetHosts(ctx)
}

func (s *HostService) GetHost(ctx context.Context, id uint) (*models.Host, error) {
	return s.hostRepo.GetHost(ctx, id)
}

func (s *HostService) GetAvailableBookingSlotsOfHost(ctx context.Context, hostId uint, start time.Time, end time.Time) ([]models.BookingSlot, error) {
	bookings, err := s.bookingRepo.GetBookingsOfHostBetween(ctx, hostId, start, end)
	if err != nil {
		return nil, err
	}

	availableSlots := make([]models.BookingSlot, 0)

	// If there are no bookings, the entire period is available
	if len(bookings) == 0 {
		availableSlots = append(availableSlots, models.BookingSlot{
			Start: start,
			End:   end,
		})
		return availableSlots, nil
	}

	// Sort bookings by check-in date
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].CheckInDate.Before(bookings[j].CheckInDate)
	})

	// Check if there's a gap between start time and first booking
	if start.Before(bookings[0].CheckInDate) {
		availableSlots = append(availableSlots, models.BookingSlot{
			Start: start,
			End:   bookings[0].CheckInDate,
		})
	}

	// Find gaps between bookings
	for i := 0; i < len(bookings)-1; i++ {
		if bookings[i].CheckOutDate.Before(bookings[i+1].CheckInDate) {
			availableSlots = append(availableSlots, models.BookingSlot{
				Start: bookings[i].CheckOutDate,
				End:   bookings[i+1].CheckInDate,
			})
		}
	}

	// Check if there's a gap between last booking and end time
	lastBooking := bookings[len(bookings)-1]
	if lastBooking.CheckOutDate.Before(end) {
		availableSlots = append(availableSlots, models.BookingSlot{
			Start: lastBooking.CheckOutDate,
			End:   end,
		})
	}

	return availableSlots, nil
}

func (s *HostService) CreateHost(ctx context.Context, host *models.Host) error {
	return s.hostRepo.CreateHost(ctx, host)
}
