package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lentscode/booking-server/internals/models"
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
	router.POST("/slots", a.GetAvailableSlotsOfHost)

	router.Run(":8080")
}

func (a *Api) SignUp(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "first_name and last_name are required"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	sessionId, err := a.userService.SignUp(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"session_id": sessionId})
}

func (a *Api) Login(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	sessionId, err := a.userService.Login(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session_id": sessionId})
}

func (a *Api) GetHosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	hosts, err := a.hostService.GetHosts(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hosts)
}

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
