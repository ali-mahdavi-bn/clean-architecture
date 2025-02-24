package adapter

import (
	"gorm.io/gorm"
	"time"
)

type Entity interface {
	GetID() uint
	IsDeleted() bool
}

type BaseEntity struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *BaseEntity) GetID() uint {
	return u.ID
}
func (u *BaseEntity) IsDeleted() bool {
	return u.DeletedAt.Valid
}
