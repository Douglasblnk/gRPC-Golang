package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	CustomerID    string         `gorm:"not null"`
	CustomerName  string         `gorm:"not null"`
	CustomerEmail string         `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"default:now();not null"`
	UpdatedAt     time.Time      `gorm:"default:now();not null"`
	DeletedAt     gorm.DeletedAt ``
	Items         []OrderItem
}

type OrderItem struct {
	ID        int    `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	OrderID   string `gorm:"not null;"`
	ProductID string `gorm:"not null"`
	Qtd       int    `gorm:"not null"`
}
