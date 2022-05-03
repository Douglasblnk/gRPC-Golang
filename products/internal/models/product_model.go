package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        string         `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	Title     string         `gorm:"not null"`
	Price     int            `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:now();not null"`
	UpdatedAt time.Time      `gorm:"default:now();not null"`
	DeletedAt gorm.DeletedAt ``
}
