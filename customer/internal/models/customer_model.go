package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:now();not null"`
	UpdatedAt time.Time      `gorm:"default:now();not null"`
	DeletedAt gorm.DeletedAt ``
}
