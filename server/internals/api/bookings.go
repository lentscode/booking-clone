package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lentscode/booking-server/internals/models"
)

func (a *Api) GetAvailableSlotsOfHost(c *gin.Context) {
	req := new(AvailableSlotsRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	slots, err := a.hostService.GetAvailableBookingSlotsOfHost(ctx, req.HostID, req.Start, req.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

func (a *Api) CreateBooking(c *gin.Context) {
	req := new(CreateBookingRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	slots, err := a.hostService.GetAvailableBookingSlotsOfHost(ctx, req.HostID, req.Start, req.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(slots) == 0 || len(slots) > 1 || slots[0].Start != req.Start || slots[0].End != req.End {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking slot is not available"})
		return
	}

	slot := slots[0]
	host, err := a.hostService.GetHost(ctx, req.HostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPrice := slot.End.Sub(slot.Start).Hours() / 24 * host.Price

	booking := &models.Booking{
		CheckInDate:  req.Start,
		CheckOutDate: req.End,
		TotalPrice:   totalPrice,
		Status:       "pending",
		UserID:       c.GetUint("user_id"),
		HostID:       req.HostID,
	}

	err = a.bookingService.CreateBooking(ctx, booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (a *Api) GetBooking(c *gin.Context) {
	bookingIdStr := c.Param("id")
	bookingId, err := strconv.ParseUint(bookingIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	booking, err := a.bookingService.GetBooking(ctx, uint(bookingId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (a *Api) GetBookingsOfUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	bookings, err := a.bookingService.GetBookingsOfUser(ctx, c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, bookings)
}
