package api

import "time"

type AvailableSlotsRequest struct {
	HostID int64     `json:"host_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}
