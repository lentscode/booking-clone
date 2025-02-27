package api

import "time"

type AvailableSlotsRequest struct {
	HostID uint     `json:"host_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

type CreateBookingRequest struct {
	HostID uint     `json:"host_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}
