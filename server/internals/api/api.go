package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lentscode/booking-server/internals/services"
)

type Api struct {
	userService    *services.UserService
	hostService    *services.HostService
	bookingService *services.BookingService
}

func NewApi(userService *services.UserService, hostService *services.HostService, bookingService *services.BookingService) *Api {
	return &Api{userService: userService, hostService: hostService, bookingService: bookingService}
}

func (a *Api) Start() {
	router := gin.Default()

	router.POST("/signup", a.SignUp)
	router.POST("/login", a.Login)

	router.GET("/hosts", a.GetHosts)
	router.POST("/hosts", a.CreateHost)

	router.POST("/slots", a.GetAvailableSlotsOfHost)

	router.GET("/bookings", a.GetBookingsOfUser)
	router.GET("/bookings/:id", a.GetBooking)

	router.Run(":8080")
}
