package model

import (
	"time"
)

type UserRole string

const (
	SuperAdmin UserRole = "SUPER_ADMIN"
	Admin      UserRole = "ADMIN"
)

type User struct {
	Id        int       `gorm:"primaryKey;autoIncrement:true"`
	Username  string    `gorm:"size:320;unique;not null"`
	Password  string    `gorm:"size:60;not null"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(6);not null"`
	Board     []Board   `gorm:"foreignKey:UserId"`
}
