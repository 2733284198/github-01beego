package routers

import (
	"bee01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//beego.NSCond(func(ctx *context.Context) bool {
		//	if ctx.Input.Domain() == "api.beego.me" {
		//		return true
		//	}
		//	return false
		//}),

		beego.NSNamespace("/userv1",
			//beego.NSRouter("/user/add", &controllers.UserController{}, "post:Doadd"),
			beego.NSRouter("/add", &controllers.UserController{}, "post:Doadd"),
		),
	)

	//beego.NSRouter("/changepassword", &UserController{}),
	//beego.Router("/user/add", &controllers.UserController{}, "post:Doadd")
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})

	beego.Router("/food", &controllers.WmfoodController{})

	// gorm
	beego.Router("/user/gwhere/:id", &controllers.UserController{}, "get:GormWhere")

	beego.Router("/user/gget", &controllers.UserController{}, "get:Gormget")
	beego.Router("/user/gadd", &controllers.UserController{}, "get:Gadd")
	beego.Router("/user/gedit/:id", &controllers.UserController{}, "get:Gedit")

	// user
	beego.Router("/user/jump", &controllers.UserController{}, "get:Jump")

	beego.Router("/user/testsession", &controllers.UserController{}, "get:Testsession")

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id", &controllers.UserController{}, "get:Getuid")
	beego.Router("/user/:id([0-9]+).html", &controllers.UserController{}, "get:Getuidhtml")

	// 添加用户
	beego.Router("/user/json", &controllers.UserController{}, "get:GetOne")
	beego.Router("/user/addform", &controllers.UserController{}, "get:AddForm")

	//beego.Router("/user/add", &controllers.UserController{}, "post:Doadd")

	// xml类型，支付宝，微信返回类型
	beego.Router("/user/xml", &controllers.UserController{}, "post:Xml")
	beego.Router("/user/pay", &controllers.UserController{}, "get:Pay")
	beego.Router("/user/paycallback", &controllers.UserController{}, "get:PayCallback")
}
