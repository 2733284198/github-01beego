package main

import (
	_ "bee01/controllers"
	_ "bee01/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()
}
