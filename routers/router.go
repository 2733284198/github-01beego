package routers

import (
	"bee01/controllers"
	"bee01/controllers/admin"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSBefore(func(ctx *context.Context) {
			//ctx.Output.Body([]byte("notAllowed"))
			fmt.Println("我是一个中间件")
			beego.Warning("我是一个中间件")

			//ctx.WriteString("/n/n v1-NSBefore")
		}),

		//beego.NSRouter("/", &controllers.NameController{}),
		beego.NSRouter("/", &admin.NameController{}),
		beego.NSRouter("/name", &admin.NameController{}),

		beego.NSRouter("/login", &admin.NameController{}, "get:Login"),

		beego.NSRouter("/dologin", &admin.NameController{}, "post:DoLogin"),

		//beego.NSRouter("/", &admin.MainController{}),
		//beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),

		beego.NSRouter("/food", &controllers.WmfoodController{}),

		//beego.NSCond(func(ctx *context.Context) bool {
		//	if ctx.Input.Domain() == "api.beego.me" {
		//		return true
		//	}
		//	return false
		//}),

		/*beego.NSNamespace("/userv1",
			beego.NSRouter("/add", &controllers.UserController{}, "post:Doadd"),
		),*/
	)

	//beego.NSRouter("/changepassword", &UserController{}),
	//beego.Router("/user/add", &controllers.UserController{}, "post:Doadd")
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/img", &controllers.MainController{}, "get:Img")

	//  图像处理
	beego.Router("/qrcode", &controllers.MainController{}, "get:Qrcode")
	beego.Router("/scaleimg", &controllers.MainController{}, "get:ScaleImg")

	//beego.Router("/food", &controllers.WmfoodController{})

	// gorm
	beego.Router("/user/gwhere/:id", &controllers.UserController{}, "get:GormWhere")

	beego.Router("/user/gget", &controllers.UserController{}, "get:Gormget")
	beego.Router("/user/gadd", &controllers.UserController{}, "get:Gadd")
	beego.Router("/user/gedit/:id", &controllers.UserController{}, "get:Gedit")

	// user
	beego.Router("/user/jump", &controllers.UserController{}, "get:Jump")

	// json: todo 2020-12-24 09:56:50
	beego.Router("/jj", &controllers.MainController{}, "get:Jj")
	beego.Router("/jj1", &controllers.MainController{}, "get:Jj1")
	beego.Router("/jj2", &controllers.MainController{}, "get:Jj2")

	beego.Router("/product", &controllers.MainController{}, "get:Product")

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
