package models

import "time"

type BookingSlot struct {
	Start time.Time
	End   time.Time
}
