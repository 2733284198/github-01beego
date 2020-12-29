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
	/*
		查询 person数据
	*/

	var person Person
	// id查询
	result, err := engine.Id(1).Get(&person)

	fmt.Println("\n ====>person ", person.ID, person.Name, person.Age)
	fmt.Println(result, err)

	// where查询
	var person1 Person
	//engine.Where("id = ?", 1).Get(&person1)
	engine.Where("age = ?", 2).Get(&person1)
	fmt.Println("\n ====>person1 ", person1.ID, person1.Name, person1.Age)

	// 不写条件，返回所有数据
	var person1all []Person
	//engine.Where("id = ?", 1).Get(&person1)
	//engine.Find(&person1all)
	engine.OrderBy("age desc").Find(&person1all)
	fmt.Println("\n ====>person1all ", person1all)

	//engine.Where("age = ?", 2).Get(&person1s)
	// Find：返回多条数据
	var person1s []Person
	engine.Where("age = ?", 2).Find(&person1s) //
	//fmt.Println("\n ====>person1s  ", person1s.ID, person1s.Name, person1s.Age)
	fmt.Println("\n ====>person1s  ", person1s)

	// and
	var person1sand []Person
	//engine.Where("age = ?", 2).And("id = ?", 4).Find(&person1s) //
	//engine.Where("age = ?", 2).And("id = ?", 4).Get(&person1) //
	engine.Where("age = ?", 2).And("id = ?", 4).Find(&person1sand) //
	fmt.Println("\n ====>person1sand  ", person1sand)
	//fmt.Println("\n ====>person1  ", person1)

	// 原生sql
	var personSql []Person
	engine.SQL("select * from person where age > 2 ").Find(&personSql)
	fmt.Println("\n ====>personSql  ", personSql)
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
