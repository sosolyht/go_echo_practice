package model

import "time"

type Media struct {
	Id         string    `gorm:"primaryKey;autoIncrement:true"`
	Title      string    `gorm:"size:100;not null"` // filename 은 aws.com/ <= 슬래쉬 뒤의 파일명으로 반환 하면 될듯?
	Size       string    `gorm:"size:50;not null"`  // 224749240 <= 이런 형식인데..어뜨케 mb 나 gb 로 변환?..
	Codec      string    `gorm:"size:20;not null"`  // codec_long_name 의 첫번째를 split 해서 [0] 번째꺼
	Duration   string    `gorm:"size:30;not null"`
	Resolution string    `gorm:"size:15;not null"` // resolution 은 width 와 height 붙여서 저장 ex) 1920 + x + 1080
	CreatedAt  time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt  time.Time `gorm:"type:datetime(6);not null"`
}

func (Media) TableName() string {
	return "medias"
}
