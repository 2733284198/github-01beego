package routers

import (
	"bee01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/food", &controllers.WmfoodController{})

	// user
	beego.Router("/user/jump", &controllers.UserController{}, "get:Jump")

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id", &controllers.UserController{}, "get:Getuid")
	beego.Router("/user/:id([0-9]+).html", &controllers.UserController{}, "get:Getuidhtml")

	beego.Router("/user/json", &controllers.UserController{}, "get:GetOne")
	beego.Router("/user/add", &controllers.UserController{}, "post:DoAdd")

	// xml类型，支付宝，微信返回类型
	beego.Router("/user/xml", &controllers.UserController{}, "post:Xml")
	beego.Router("/user/pay", &controllers.UserController{}, "get:Pay")
}
