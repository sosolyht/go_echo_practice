package database

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"time"
)

var a struct {
	DB struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Host string `json:"host"`
		Port int    `json:"port"`
		Name string `json:"name"`
	} `json:"db"`
	IsDebug bool `json:"is_debug"`
}

const mysqlData = "%s:%s@tcp(%s:%d)/%s?%s"

func Connect() {
	file, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}
	var val = make(url.Values)
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "true")
	val.Add("loc", time.UTC.String())

	err = json.NewDecoder(file).Decode(&a)

	data := a.DB
	DBConn := fmt.Sprintf(mysqlData,
		data.User, data.Pass, data.Host, data.Port, data.Name, val.Encode())

	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetMaxOpenConns(15)
	return
}
