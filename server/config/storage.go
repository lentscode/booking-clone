package config

import (
	"log"
	"os"

	"github.com/lentscode/booking-server/internals/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

func NewStorage() *Storage {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")))

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Host{}, &models.Booking{})

	return &Storage{db}
}
