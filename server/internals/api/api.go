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

	router.POST("/users", a.CreateUser)
	router.GET("/hosts", a.GetHosts)

	router.Run(":8080")
}

func (a *Api) CreateUser(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := a.userService.CreateUser(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
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
