package core

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"
)

var c struct {
	DB struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Host string `json:"host"`
		Port uint16 `json:"port"`
		Name string `json:"name"`
	} `json:"db"`

	IsDebug bool `json:"is_debug"`
}

var (
	IsDebug = true
	DBConn  = ""
)

const (
	mysqlDBConnFormat = "%s:%s@tcp(%s:%d)/%s?%s"
)

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	var val = make(url.Values)
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "true")
	val.Add("loc", time.UTC.String())

	err = json.NewDecoder(file).Decode(&c)

	if err != nil {
		IsDebug = true
		DBConn = fmt.Sprintf(mysqlDBConnFormat,
			"root", "1234", "localhost", 3306, "editfolio", val.Encode())
	} else {
		var db = c.DB

		IsDebug = c.IsDebug
		DBConn = fmt.Sprintf(mysqlDBConnFormat,
			db.User, db.Pass, db.Host, db.Port, db.Name, val.Encode())
	}
}
