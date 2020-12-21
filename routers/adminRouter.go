package routers

import (
	//"bee01/controllers"
	"bee01/controllers/admin"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns :=
		beego.NewNamespace("/admin",
			//中间件:匹配路由前会执,可以用于权限验证
			//注意引入的包： github.com/astaxie/beego/context
			beego.NSBefore(func(ctx *context.Context) {
				fmt.Println("我是一个中间件，匹配路由之前执行")
			}),
			beego.NSRouter("/", &admin.MainController{}),
			beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),

			beego.NSRouter("/manager", &admin.ManagerController{}),
			beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
			beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),

			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/focus", &admin.FocusController{}),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
