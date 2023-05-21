package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	UserID uint
	Status string
	Reason string
}
