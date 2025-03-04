package model

import "clean-hex/internal/user_management/domain/entities"

type Trade struct {
	entities.Trade
	UserID uint   `gorm:"index"`
	Stock  string `gorm:"not null"`
	Price  int    `gorm:"not null"`
	Amount int    `gorm:"not null"`
}
