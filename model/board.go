package model

import "github.com/google/uuid"

type Board struct {
	Id      uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title   string    `gorm:"size:30;not null"`
	Content string    `gorm:"size:1000;not null"`
	UserId  uuid.UUID
}
