package service

import (
	"ginstudy/logger"
	"ginstudy/model"

	"github.com/jinzhu/gorm"
)

func init() {
	var mysql *gorm.DB
	var err error
	if mysql, err = gorm.Open("mysql", "root:my-secret-pw@tcp(101.43.71.14:3306)/ginstudy?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
		logger.Errorf("connection database error, %v", err)
	} else if nil != mysql {
		defer func(mysql *gorm.DB) {
			if err := mysql.Close(); nil != err {
				logger.Errorf("Disconnection from database failed, %v", err)
			}
		}(mysql)
		logger.Debug("connect success")

		// 需要自动创建表的模型
		if err = mysql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Product{}).Error; nil != err {
			logger.Errorf("auto migrate tables failed: %v", err)
		}

	}
}
