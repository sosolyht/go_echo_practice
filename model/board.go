package model

import (
	"time"
)

type Board struct {
	Id        int       `gorm:"primaryKey;autoIncrement:true"`
	Title     string    `gorm:"size:30;not null"`
	Content   string    `gorm:"size:1000;not null"`
	UserId    int       `gorm:"type:char(36)"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(6);not null"`
}
