package service

import (
	"ginstudy/logger"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var once sync.Once

func Init() {
	once.Do(func() {
		var err error
		if db, err = gorm.Open("mysql", "root:my-secret-pw@tcp(101.43.71.14:3306)/ginstudy?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
			logger.Errorf("connection database error, %v", err)
		} else {
			logger.Debug("connect success!")
			db.DB().SetMaxIdleConns(10)
			db.DB().SetMaxOpenConns(50)
			db.LogMode(true)
		}
	})
}

func Close() {
	if db != nil {
		if err := db.Close(); nil != err {
			logger.Errorf("Disconnect from database failed: %v", err)
		}
	}
	db = nil
}
