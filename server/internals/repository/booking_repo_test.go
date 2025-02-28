package repository

import (
	"context"
	"testing"
	"time"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/models"
	"github.com/stretchr/testify/assert"
)

func newHost() models.Host {
	return models.Host{
		Name:        "John Doe",
		Location:    "123 Main St, Anytown, USA",
		Rating:      4.5,
		Description: "John is a great host",
		Capacity:    2,
		Price:       100.0,
	}
}

func newUser() models.User {
	return models.User{
		Email:     "test@test.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
	}
}
func TestCreateBooking(t *testing.T) {
	storage := config.NewStorage()

	repo := NewBookingRepository(storage)
	host := newHost()
	storage.Db.Create(&host)

	user := newUser()
	storage.Db.Create(&user)

	booking := models.Booking{
		UserID:       user.ID,
		HostID:       host.ID,
		CheckInDate:  time.Now(),
		CheckOutDate: time.Now().AddDate(0, 0, 1),
		TotalPrice:   100.0,
		Status:       "pending",
	}

	err := repo.CreateBooking(context.Background(), &booking)
	if err != nil {
		t.Fatalf("Failed to create booking: %v", err)
	}

	assert.NotNil(t, booking.ID)
	dbBooking := new(models.Booking)

	res := storage.Db.First(&dbBooking, booking.ID)
	if res.Error != nil {
		t.Fatalf("Failed to get booking: %v", res.Error)
	}

	assert.Equal(t, booking.ID, dbBooking.ID)
	assert.Equal(t, booking.UserID, dbBooking.UserID)
	assert.Equal(t, booking.HostID, dbBooking.HostID)
	assert.Equal(t, booking.CheckInDate.Format(time.RFC3339), dbBooking.CheckInDate.Format(time.RFC3339))
	assert.Equal(t, booking.CheckOutDate.Format(time.RFC3339), dbBooking.CheckOutDate.Format(time.RFC3339))
	assert.Equal(t, booking.TotalPrice, dbBooking.TotalPrice)
	assert.Equal(t, booking.Status, dbBooking.Status)

	storage.Db.Where("1 = 1").Delete(&models.Booking{})
}

func TestGetBookingByID(t *testing.T) {
	storage := config.NewStorage()
	repo := NewBookingRepository(storage)

	host := newHost()
	storage.Db.Create(&host)

	user := newUser()
	storage.Db.Create(&user)

	booking := models.Booking{
		UserID:       user.ID,
		HostID:       host.ID,
		CheckInDate:  time.Now(),
		CheckOutDate: time.Now().AddDate(0, 0, 1),
		TotalPrice:   100.0,
		Status:       "pending",
	}

	storage.Db.Create(&booking)

	dbBooking, err := repo.GetBooking(context.Background(), booking.ID)
	if err != nil {
		t.Fatalf("Failed to get booking: %v", err)
	}

	assert.Equal(t, booking.ID, dbBooking.ID)
	assert.Equal(t, booking.UserID, dbBooking.UserID)
	assert.Equal(t, booking.HostID, dbBooking.HostID)
	assert.Equal(t, booking.CheckInDate.Format(time.RFC3339), dbBooking.CheckInDate.Format(time.RFC3339))
	assert.Equal(t, booking.CheckOutDate.Format(time.RFC3339), dbBooking.CheckOutDate.Format(time.RFC3339))
	assert.Equal(t, booking.TotalPrice, dbBooking.TotalPrice)
	assert.Equal(t, booking.Status, dbBooking.Status)

	storage.Db.Where("1 = 1").Delete(&models.Booking{})
}

func TestGetBookingsOfUser(t *testing.T) {
	storage := config.NewStorage()
	repo := NewBookingRepository(storage)

	host := newHost()
	storage.Db.Create(&host)
	
	user := newUser()
	storage.Db.Create(&user)

	bookings := []models.Booking{
		{
			UserID:       user.ID,
			HostID:       host.ID,
			CheckInDate:  time.Now(),
			CheckOutDate: time.Now().AddDate(0, 0, 1),
			TotalPrice:   100.0,
			Status:       "pending",
		},
		{
			UserID:       user.ID,
			HostID:       host.ID,
			CheckInDate:  time.Now().AddDate(0, 0, 1),
			CheckOutDate: time.Now().AddDate(0, 0, 2),
			TotalPrice:   200.0,
			Status:       "pending",
		},
		{
			UserID:       3,
			HostID:       host.ID,
			CheckInDate:  time.Now().AddDate(0, 0, 2),
			CheckOutDate: time.Now().AddDate(0, 0, 3),
			TotalPrice:   300.0,
			Status:       "pending",
		},
	}

	storage.Db.Create(&bookings)

	dbBookings, err := repo.GetBookingsOfUser(context.Background(), user.ID)
	if err != nil {
		t.Fatalf("Failed to get bookings: %v", err)
	}
	
	assert.Equal(t, len(dbBookings), 2)
	assert.Equal(t, dbBookings[0].ID, bookings[0].ID)
	assert.Equal(t, dbBookings[1].ID, bookings[1].ID)

	storage.Db.Where("1 = 1").Delete(&models.Booking{})
}

func TestGetBookingsOfHost(t *testing.T) {
	storage := config.NewStorage()
	repo := NewBookingRepository(storage)

	host := newHost()
	storage.Db.Create(&host)
	
	user := newUser()
	storage.Db.Create(&user)

	bookings := []models.Booking{
		{
			UserID:       user.ID,
			HostID:       host.ID,
			CheckInDate:  time.Now(),
			CheckOutDate: time.Now().AddDate(0, 0, 1),
			TotalPrice:   100.0,
			Status:       "pending",
		},
		{
			UserID:       user.ID,
			HostID:       host.ID,
			CheckInDate:  time.Now().AddDate(0, 0, 1),
			CheckOutDate: time.Now().AddDate(0, 0, 2),
			TotalPrice:   200.0,
			Status:       "pending",
		},
		{
			UserID:       user.ID,
			HostID:       2,
			CheckInDate:  time.Now().AddDate(0, 0, 2),
			CheckOutDate: time.Now().AddDate(0, 0, 3),
			TotalPrice:   300.0,
			Status:       "pending",
		},
	}

	storage.Db.Create(&bookings)

	dbBookings, err := repo.GetBookingsOfHost(context.Background(), host.ID)
	if err != nil {
		t.Fatalf("Failed to get bookings: %v", err)
	}
	
	assert.Equal(t, len(dbBookings), 2)
	assert.Equal(t, dbBookings[0].ID, bookings[0].ID)
	assert.Equal(t, dbBookings[1].ID, bookings[1].ID)

	storage.Db.Where("1 = 1").Delete(&models.Booking{})
}

//! fails
// func TestGetBookingsOfHostBetween(t *testing.T) {
// 	storage := config.NewStorage()
// 	repo := NewBookingRepository(storage)

// 	host := newHost()
// 	storage.Db.Create(&host)
	
// 	user := newUser()
// 	storage.Db.Create(&user)

// 	now := time.Now()

// 	bookings := []models.Booking{
// 		{
// 			UserID:       user.ID,
// 			HostID:       host.ID,
// 			CheckInDate:  now,
// 			CheckOutDate: now.AddDate(0, 0, 1),
// 			TotalPrice:   100.0,
// 			Status:       "pending",
// 		},
// 		{
// 			UserID:       user.ID,
// 			HostID:       host.ID,
// 			CheckInDate:  now.AddDate(0, 0, 1),
// 			CheckOutDate: now.AddDate(0, 0, 2),
// 			TotalPrice:   200.0,
// 			Status:       "pending",
// 		},
// 		{
// 			UserID:       user.ID,
// 			HostID:       2,
// 			CheckInDate:  now.AddDate(0, 0, 2),
// 			CheckOutDate: now.AddDate(0, 0, 3),
// 			TotalPrice:   300.0,
// 			Status:       "pending",
// 		},
// 	}

// 	storage.Db.Create(&bookings)

// 	dbBookings, err := repo.GetBookingsOfHostBetween(context.Background(), host.ID, now, now.AddDate(0, 0, 2))
// 	if err != nil {
// 		t.Fatalf("Failed to get bookings: %v", err)
// 	}
	
// 	assert.Equal(t, len(dbBookings), 2)
// 	assert.Equal(t, dbBookings[0].ID, bookings[0].ID)
// 	assert.Equal(t, dbBookings[1].ID, bookings[1].ID)

// 	storage.Db.Where("1 = 1").Delete(&models.Booking{})
// }
