package database

import (
	"go_echo/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (db *gorm.DB) {
	var logLevel = logger.Info

	if !core.IsDebug {
		logLevel = logger.Warn
	}

	db, err := gorm.Open(mysql.Open(core.DBConn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
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
