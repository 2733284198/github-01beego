package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"xorm.io/core"
	//"github.com/go-xorm/core"
)

func main() {

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
	engine.Sync(new(Person))

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

type Person struct {
	Age  int
	Name string
}

/**
 *
 */
func OrmMapping() {

}
