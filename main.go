package main

import (
	"bee01/controllers"
	_ "bee01/controllers"
	_ "bee01/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Router("/user", &controllers.MainController{})
	beego.Run()
}
