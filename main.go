package main

import (
	"bee01/models"
	_ "bee01/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"
	//_ "bee01/controllers"
	_ "bee01/routers"
	"encoding/gob"
)

func init() {
	fmt.Println("main-init")

	gob.Register(models.User{})
}

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

func TestModel() {
	_ = models.Db.AutoMigrate(models.Article{})
}

func TestLog() {
	//beego.SetLogger("")
	logs.Alert("====> Logs modules ")

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	// 不打印在console，只打印在文件
	//beego.BeeLogger.DelLogger("console")

	// beego
	beego.Info("这是info")
	beego.Debug("这是 debug")
	beego.Critical("这是 Critical")
	beego.Alert("这是 Alert")

	// 只打印debug信息
	beego.SetLevel(beego.LevelDebug)
}

func main() {

	//beego.Router("/user", &controllers.MainController{})
	//beego.Router("/food", &controllers.WmfoodController{})

	//getconfig()
	//setsession()

	TestLog()
	TestModel()
	beego.Run(":88")

	//beego.Run()
	// 关闭数据库连接
	//defer models.DB.Close()
}
