package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(127.0.0.1:3306)/beego?charset=utf8mb4&parseTime=True&loc=Local"

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
