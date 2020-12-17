package main

import (
	_ "bee01/models"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
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

	// dev prod
	devhttpport := beego.AppConfig.String("dev::httpport")
	beego.Info(devhttpport)

	prodhttpport := beego.AppConfig.String("prod::httpport")
	beego.Info(prodhttpport)

	fmt.Println(port)
}

func setsession() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
}

func main() {

	//beego.Router("/user", &controllers.MainController{})
	//beego.Router("/food", &controllers.WmfoodController{})

	getconfig()
	setsession()

	beego.Run(":88")

	//beego.Run()
	// 关闭数据库连接
	//defer models.DB.Close()
}
