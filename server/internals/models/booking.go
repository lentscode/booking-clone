package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	CheckInDate  time.Time `json:"check_in_date" binding:"required"`
	CheckOutDate time.Time `json:"check_out_date" binding:"required"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`

	UserID uint `json:"user_id" binding:"required"`
	User   User

	HostID uint `json:"host_id" binding:"required"`
	Host   Host
}

type BookingSlot struct {
	Start time.Time
	End   time.Time
}
