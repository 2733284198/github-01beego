package models

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(127.0.0.1:3306)/beego?charset=utf8mb4&parseTime=True&loc=Local"

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		//err = beego.Error()
		fmt.Println(err)
	}

}
