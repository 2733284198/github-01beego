package models

import (
	"fmt"
	"os"
	"time"

	"gitea.com/lunny/log"
	_ "github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,   // 慢 SQL 阈值
		LogLevel:      logger.Silent, // Log level
		Colorful:      false,         // 禁用彩色打印
	},
)

func init() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
		Logger: newLogger,
	})

	if err != nil {
		//err = beego.Error()
		fmt.Println(err)
	}
}
