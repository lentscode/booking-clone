package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
