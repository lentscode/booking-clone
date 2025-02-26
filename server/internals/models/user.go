package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Sessions []UserSession
}
