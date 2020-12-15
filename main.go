package main

import (
	"github.com/astaxie/beego"

	//_ "bee01/controllers"
	_ "bee01/routers"
)

func main() {

	//beego.Router("/user", &controllers.MainController{})
	//beego.Router("/food", &controllers.WmfoodController{})

	beego.Run(":8089")

}
