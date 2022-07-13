package model

import "github.com/google/uuid"

type Board struct {
	Id      int       `gorm:"primaryKey;autoIncrement:true"`
	Title   string    `gorm:"size:30;not null"`
	Content string    `gorm:"size:1000;not null"`
	UserId  uuid.UUID `gorm:"type:char(36)"`
}
