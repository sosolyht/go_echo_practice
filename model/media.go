package model

import "time"

type Media struct {
	Id         string    `gorm:"primaryKey;autoIncrement:true"`
	Size       string    `gorm:"size:50;not null"`
	Codec      string    `gorm:"size:20;not null"`
	Duration   string    `gorm:"size:30;not null"`
	Resolution string    `gorm:"size:15;not null"`
	CreatedAt  time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt  time.Time `gorm:"type:datetime(6);not null"`
}

func (Media) TableName() string {
	return "medias"
}
