package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"xorm.io/core"
	//"github.com/go-xorm/core"
)

func getOrmData() {
	//var person *Person
	var person Person

	bool, err := engine.Id(1).Get(&person)
	if err != nil {
		fmt.Println(err.Error())
	}

	if bool == true {
		fmt.Println("Person is true ")
	} else {
		fmt.Println("Person is false ")
	}

	fmt.Println(person.ID, person.Name, person.Age)
}

var engine *xorm.Engine

func initData() {

	//1. 创建数据库引擎对象
	//engine, err := xorm.NewEngine("mysql", "root:yu271400@/elmcms?charset=utf8")
	engine, err := xorm.NewEngine("mysql", "root:root@/beego?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	//2. 数据库引擎关闭
	defer engine.Close()

	//数据库引擎设置

	engine.ShowSQL(true)                     //设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别
	engine.SetMaxOpenConns(10)               //设置最大连接数

	//engine.SetMaxIdleConns(2)
	engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(Person))

	// 查询 person数据

	var person Person
	engine.Id(1).Get(&person)

	fmt.Println("\n ====> ", person.ID, person.Name, person.Age)

}

func getAllData() {
	//查询表的所有数据
	//session := engine.Table("user")
	session := engine.Table("person")
	count, err := session.Count()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(count)

	//使用原生sql语句进行查询
	//result, err := engine.Query("select  * from user")
	result, err := engine.Query("select  * from person")
	if err != nil {
		panic(err.Error())
	}
	for key, value := range result {
		fmt.Println(key, value)
	}
}

func main() {
	initData()

	//getOrmData()
}

type Person struct {
	ID   int `xorm:"pk autoincr" `
	Age  int
	Name string
}

/**
 *
 */
func OrmMapping() {

}
