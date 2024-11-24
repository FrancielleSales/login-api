package models

import (
	"time"
)

type Users struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
	Email string `gorm:"unique;size:100;not null"`
	Password string `gorm:"size:80;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}