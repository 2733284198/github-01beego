package routers

import (
	"bee01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/food", &controllers.WmfoodController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/json", &controllers.UserController{}, "get:GetOne")
	beego.Router("/user/add", &controllers.UserController{}, "post:DoAdd")

	// xml类型，支付宝，微信返回类型
	beego.Router("/user/xml", &controllers.UserController{}, "post:Xml")
	beego.Router("/user/pay", &controllers.UserController{}, "get:Pay")
}
