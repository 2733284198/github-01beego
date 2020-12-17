package main

import (
	"fmt"
	"github.com/astaxie/beego"
	//_ "bee01/controllers"
	_ "bee01/routers"
)

func getconfig() {
	port := beego.AppConfig.String("httpport")
	beego.Info(port)

	admins := beego.AppConfig.Strings("admin")
	beego.Info(admins)

	beego.AppConfig.Set("rediskey", "redisk1")

	rediskey1 := beego.AppConfig.Strings("rediskey")
	beego.Info(rediskey1)

	fmt.Println(port)
}

func main() {

	//beego.Router("/user", &controllers.MainController{})
	//beego.Router("/food", &controllers.WmfoodController{})

	getconfig()

	beego.Run(":8089")
	//beego.Run()

}
