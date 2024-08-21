package core

import (
	"bubblepop_master/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Gorm() *gorm.DB {
	return MysqlConnect()
}

func MysqlConnect() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		logrus.Warn("未配置mysql,取消gorm注接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	if global.Config.System.Env == "debug" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		logrus.Warn(fmt.Sprintf("%s mysql连接失败", dsn))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	return db
}
