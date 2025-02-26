package main

import (
	"log"

	"github.com/lentscode/booking-server/config"
	"github.com/lentscode/booking-server/internals/api"
	"github.com/lentscode/booking-server/internals/repository"
	"github.com/lentscode/booking-server/internals/services"
)

func main() {
	storage := config.NewStorage()

	log.Println("Starting server...")

	userRepo := repository.NewUserRepository(storage)
	hostRepo := repository.NewHostRepository(storage)
	bookingRepo := repository.NewBookingRepository(storage)

	userService := services.NewUserService(userRepo)
	hostService := services.NewHostService(hostRepo, bookingRepo)
	bookingService := services.NewBookingService(bookingRepo)

	api := api.NewApi(userService, hostService, bookingService)
	api.Start()
}
