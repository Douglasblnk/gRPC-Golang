package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ProductID int `gorm:"not null"`
	Quantity  int `gorm:"not null"`
}

type Order struct {
	ID         int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	CustomerID int            `gorm:"not null"`
	Cart       []Cart         ``
	CreatedAt  time.Time      `gorm:"default:now();not null"`
	UpdatedAt  time.Time      `gorm:"default:now();not null"`
	DeletedAt  gorm.DeletedAt ``
}
