package models

import "gorm.io/gorm"

type UserSession struct {
	gorm.Model
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
}
