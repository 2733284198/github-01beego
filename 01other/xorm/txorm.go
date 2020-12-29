package main

import (
	//"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Users struct {
	Id       int64
	Username string
	Password int64
	Age      int64
}

type Articles struct {
	Id      int64
	Title   string
	Content string
}

//var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@/beego?charset=utf8")
	if err != nil {
		println(err)
	}

	//addlog()

	sql := "select * from users"
	results, err := engine.Query(sql)

	fmt.Println(results)

	// 生成表
	//err := engine1.Sync2(new(Articles))
	//err := engine1.Sync(new(Users))

	var users []Users
	//err := engine1.Where("name = ?", name).And("age > 10").Limit(10, 0).Find(&users)

	//err := engine1.Find(&users)

	//fmt.Println(users)

	//engine1.TableName("Users").find()
}

func addlog() {
	/* 日志1 */
	/*f, err := os.Create("./sql.log")
	if err != nil {
		println(err.Error())
		return
	}

	engine1.SetLogger(xorm.NewSimpleLogger(f))*/

	/*日志2*/
	/*logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.Fatalf("Fail to create xorm system logger: %v\n", err)
	}

	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	engine1.SetLogger(logger)*/

}
