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
	"github.com/gomodule/redigo/redis"
	//_ "github.com/go-redis/redis"
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

func TRedis() {
	conn, _ := redis.Dial("tcp", ":6379")

	_ = conn.Send("set", "name1", "lisi1")
	_ = conn.Send("mset", "age1", 11, "score", 101)

	_ = conn.Flush()

	reply, err := conn.Receive()
	if err != nil {
		logs.Warning("redis 设置内容错误")
	}

	logs.Info("===> redis ok")
	logs.Info(reply)

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

func tmap() {
	/**
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	var tags map[string]interface{}
	tags = make(map[string]interface{})

	var tagsLocal map[string]string
	tagsLocal = make(map[string]string)
	var tagsTest map[string]string
	tagsTest = make(map[string]string)
	var tagsProduction map[string]string
	tagsProduction = make(map[string]string)

	tagsLocal["dev.www.9178.us"] = "dev.www.9178.us"
	tagsLocal["dev.static.9178.us"] = "dev.static.9178.us"

	tagsTest["dev.www.9178.us"] = "www.ninja911.com"
	tagsTest["dev.static.9178.us"] = "static.ninja911.com"

	tagsProduction["dev.www.9178.us"] = "ipx.www.ninja911.com"
	tagsProduction["dev.static.9178.us"] = "ipx.static.ninja911.com"

	tags["local"] = tagsLocal
	tags["test"] = tagsLocal
	tags["production"] = tagsLocal

	fmt.Println(tags)
}

func main() {

	//beego.Router("/user", &controllers.MainController{})
	//beego.Router("/food", &controllers.WmfoodController{})

	//getconfig()
	//setsession()

	TestLog()
	TestModel()
	tmap()
	TRedis()
	beego.Run(":88")

	//beego.Run()
	// 关闭数据库连接
	//defer models.DB.Close()
}
