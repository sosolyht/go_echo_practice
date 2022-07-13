package config

import (
	"encoding/json"
	"fmt"
	"go_echo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
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

func DBConnection() {
	file, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}

	json.NewDecoder(file).Decode(&a)

	data := a.DataBase
	DBConn := fmt.Sprintf(mysqlData,
		data.User, data.Pass, data.Host, data.Port, data.Name)

	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	migrateError := db.AutoMigrate(
		&model.User{},
		&model.Board{})

	if migrateError != nil {
		panic(migrateError)
	}

	return
}
