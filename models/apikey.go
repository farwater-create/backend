package models

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Key         string `gorm:"unique;not null"`
	Permissions string `gorm:"not null"`
	UserID      uint
}
