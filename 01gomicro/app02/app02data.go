package main

import (
	"fmt"
	//"os"

	"github.com/douyu/jupiter/pkg/store/gorm"
	"github.com/douyu/jupiter/pkg/xlog"
)

type User2 struct {
	//Id   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func main() {
	gormDB := gorm.StdConfig("test").Build()
	fmt.Println(gormDB)

	//os.Exit()

	//models := []interface{}{
	//	&User{},
	//}
	//gormDB.SingularTable(true)
	//gormDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)

	//gormDB.Create(&User2{
	//Id:10
	//Name: "jupiter",
	//})

	var user User2
	gormDB.Where("id = 1").Find(&user)

	xlog.Info("user info", xlog.String("name", user.Name))
}
