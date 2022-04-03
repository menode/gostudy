package service

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var once sync.Once

func Init() {
	once.Do(func() {
		var err error
		dsn := "root:my-secret-pw@tcp(101.43.71.14:3306)/ginstudy?charset=utf8mb4&parseTime=True"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}
	})
}

// func Close() {
// 	if db != nil {
// 		if err := db.Distinct(); nil != err {
// 			logger.Errorf("Disconnect from database failed: %v", err)
// 		}
// 	}
// 	db = nil
// }
