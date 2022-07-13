package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Username  string    `gorm:"size:320;unique;not null"`
	Password  string    `gorm:"size:60;not null"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(6);not null"`
	Board     []Board   `gorm:"foreignKey:UserId"`
}
