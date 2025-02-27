package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Location    string  `json:"location" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
	Description string  `json:"description"`
	Capacity    int     `json:"capacity" binding:"required"`
	Price       float64     `json:"price" binding:"required"`
}
