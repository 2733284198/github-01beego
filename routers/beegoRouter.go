package routers

import (
	"bee01/controllers/beegoctl"
	"github.com/astaxie/beego"
)

func init() {

	ns :=
		beego.NewNamespace("/article",
			beego.NSRouter("/", &beegoctl.ArticlesController{}),
		)
	//注册 namespace
	beego.AddNamespace(ns)

	//beego.Router("/login", &index.ArticlesController{})
	//beego.Router("/bee", &beegoctl.ArticlesController{} )
}
