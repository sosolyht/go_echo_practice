package config

import (
	"encoding/json"
	"fmt"
	"go_echo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var a struct {
	DataBase struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Host string `json:"host"`
		Port int    `json:"port"`
		Name string `json:"name"`
	} `json:"db"`
	IsDebug bool `json:"is_debug"`
}

const mysqlData = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func DBConnection() *gorm.DB {
	// config.json 파일 오픈
	file, err := os.Open("config.json")

	// 에러 발생 시 panic
	if err != nil {
		panic(err)
	}

	// Json 문자열을 디코딩하여 Go 벨류로 변경
	json.NewDecoder(file).Decode(&a)

	var data = a.DataBase

	var DBConn = fmt.Sprintf(mysqlData,
		data.User, data.Pass, data.Host, data.Port, data.Name)

	var newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	// gorm 의 mysql 드라이버를 이용해 DBConn 으로 연결
	// Config 의 Logger 를 통한 쿼리 로그
	//db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})

	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// 모델 마이그레이션
	migrateError := db.AutoMigrate(
		&model.User{},
		&model.Board{})

	if migrateError != nil {
		panic(migrateError)
		fmt.Println(migrateError)
	}
	return db
}
