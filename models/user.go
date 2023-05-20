package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	DiscordID     string    `gorm:"unique;not null"`
	MinecraftUUID string    `gorm:"unique;not null"`
	Birthday      time.Time `gorm:"not null"`
	ApiKeys       []ApiKey
}

type PostUserInput struct {
	DiscordID     string    `json:"discordID" validate:"required"`
	MinecraftUUID string    `json:"minecraftUUID" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
}
